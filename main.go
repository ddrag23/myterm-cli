package main

import (
	"termius-cli/cmd"

	"github.com/spf13/cobra"
)

func main() {
	command := &cobra.Command{
		Use:   "termius-cli",
		Short: "simple ssh access",
	}
	command.AddCommand(cmd.AddSetting(), cmd.RunSsh(), cmd.ListSetting())
	command.Execute()
}
