package cmd

import (
	"github.com/Reljod/getitdone/internal/list"
	"github.com/spf13/cobra"
)

type ListCommand struct{}

func (cmd ListCommand) Create() *cobra.Command {
	return &cobra.Command{
		Use:   "ls",
		Short: "List all custom commands",
		Run: func(cmd *cobra.Command, args []string) {
			list.ListCmds()
		},
	}
}
