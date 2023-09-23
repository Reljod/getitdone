package cmd

import (
	"fmt"
	"os"

	"github.com/Reljod/getitdone/cmd/initial"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "getitdone",
	Short: "getitdone is a cli tool for saving your most used cli commands",
}

func init() {

	cmds := []*cobra.Command{
		saveCmd,
		CreateSetupCmd(),
		AddListCmd(),
	}

	rootCmd.AddCommand(cmds...)
	cmdsFromConfig := initial.CompileCommandsFromConfig()
	if cmdsFromConfig != nil {
		rootCmd.AddCommand(*cmdsFromConfig...)
	}

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
