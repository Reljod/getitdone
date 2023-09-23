package initial

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/Reljod/getitdone/internal/config"
	"github.com/spf13/cobra"
)

func CompileCommandsFromConfig() *[]*cobra.Command {
	configPath := config.ConfigPath

	f, err := os.Open(configPath)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	b, err := io.ReadAll(f)
	commandsRaw := strings.TrimSpace(string(b))
	commandsRawList := strings.Split(commandsRaw, "\n")

	if commandsRaw == "" {
		return nil
	}

	var commands []*cobra.Command = make([]*cobra.Command, 0)
	for _, c := range commandsRawList {
		splitCmd := strings.Split(c, " => ")
		name := splitCmd[0]
		command := splitCmd[1]

		cmd := &cobra.Command{
			Use:   name,
			Short: "Custom command",
			Run: func(cmd *cobra.Command, args []string) {
				out, err := exec.Command("bash", "-c", command).Output()
				if err != nil {
					log.Fatal(err)
				}

				fmt.Printf("Ran via getitdone:\n\n%v", string(out))
			},
		}

		commands = append(commands, cmd)
	}

	return &commands
}
