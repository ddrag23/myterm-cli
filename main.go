package main

import (
	"log"
	"os"
	"termius-cli/cmd"

	"github.com/spf13/cobra"
)

func main() {
	command := &cobra.Command{
		Use:   "termius-cli",
		Short: "simple ssh access",
	}
	command.AddCommand(cmd.AddSetting(), cmd.RunSsh(), cmd.ListSetting(),cmd.DeleteSetting())
	err := command.Execute()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
