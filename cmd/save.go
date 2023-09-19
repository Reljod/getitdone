package cmd

import (
	"strings"

	"github.com/Reljod/getitdone/internal/save"
	"github.com/spf13/cobra"
)

var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Save a command",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		command := strings.Join(args[1:], " ")

		save.Save(name, &command)
	},
}
