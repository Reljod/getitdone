package cmd

import (
	"fmt"
	"os"

	"github.com/Reljod/getitdone/cmd/initial"
	"github.com/Reljod/getitdone/internal/delete"
	"github.com/Reljod/getitdone/internal/list"
	"github.com/Reljod/getitdone/internal/read"
	"github.com/Reljod/getitdone/internal/save"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "getitdone",
	Short: "getitdone is a cli tool for saving your most used cli commands",
}

func init() {

	rd := read.Read{}
	del := delete.Delete{Read: &rd}
	ls := list.List{Read: &rd}
	sv := save.Save{}

	rootCmd.AddCommand(DeleteCommand{Delete: &del}.Create())
	rootCmd.AddCommand(ListCommand{List: &ls}.Create())
	rootCmd.AddCommand(SaveCommand{Save: &sv}.Create())
	rootCmd.AddCommand(SetupCommand{}.Create())
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
