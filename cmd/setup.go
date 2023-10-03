package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const defaultDirectory string = "$HOME/.getitdone"

type SetupCommand struct{}

func (setup SetupCommand) createConfigDirectory(dir string) (string, error) {
	fmt.Printf("Creating %v directory...\n", dir)

	out, err := exec.Command("bash", "-c", fmt.Sprintf("echo %v", dir)).Output()
	if err != nil {
		return "", err
	}

	dirOut := regexp.MustCompile(`[^a-zA-Z0-9\./\\]+`).ReplaceAllString(string(out), "")
	if err := os.Mkdir(dirOut, 0755); err != nil {
		return "", err
	}

	fmt.Printf("Successfully created %v directory", dirOut)
	return dirOut, nil
}

func (setup SetupCommand) run(cmd *cobra.Command, args []string) {
	directory := viper.GetString("directory")

	validate := func(input string) error {
		if input == "" {
			return nil
		}

		if _, err := os.Stat(input); os.IsNotExist(err) {
			return err
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    fmt.Sprintf("Save configurations to (default:%v)", directory),
		Validate: validate,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("%v is not valid\n", err)
		return
	}

	var inputDirectory string = directory
	if result != "" {
		inputDirectory = result
	}

	_, err = setup.createConfigDirectory(inputDirectory)
	if err != nil {
		fmt.Printf("Cannot create directory %v", err)
		return
	}
}

func (setup SetupCommand) Create() *cobra.Command {

	setupCmd := &cobra.Command{
		Use:   "setup",
		Short: "Setup getitdone",
		Run:   setup.run,
	}

	setupCmd.PersistentFlags().StringP("directory", "d", defaultDirectory,
		"Directory to save getitdone configurations and commands")
	viper.BindPFlag("directory", setupCmd.PersistentFlags().Lookup("directory"))
	return setupCmd
}
