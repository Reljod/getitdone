package cmd

import (
	"strings"

	"github.com/Reljod/getitdone/internal/save"
	"github.com/spf13/cobra"
)

type SaveCommand struct {
	save *save.Save
}

func (saveCmd SaveCommand) Create() *cobra.Command {
	return &cobra.Command{
		Use:   "save",
		Short: "Save a command",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			command := strings.Join(args[1:], " ")

			saveCmd.save.Save(name, &command)
		},
	}
}
