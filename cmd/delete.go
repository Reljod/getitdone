package cmd

import (
	"github.com/Reljod/getitdone/internal/delete"
	"github.com/spf13/cobra"
)

type DeleteCommand struct {
	Delete *delete.Delete
}

func (deleteCmd DeleteCommand) Create() *cobra.Command {
	return &cobra.Command{
		Use:   "delete",
		Short: "Delete a command",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			deleteCmd.Delete.DeleteCmd(name)
		},
	}
}
