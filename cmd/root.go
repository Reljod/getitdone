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
	rootCmd.AddCommand(saveCmd)
	rootCmd.AddCommand(CreateSetupCmd())
	rootCmd.AddCommand(*initial.CompileCommandsFromConfig()...)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
