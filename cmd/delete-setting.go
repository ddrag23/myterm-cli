package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func DeleteSetting() *cobra.Command{
	return &cobra.Command{
		Use: "delete",
		Short: "Delete your ssh config",
		Run: func(cmd *cobra.Command, args []string) {
			answer := getAllCredentials()
			err := os.Remove(fmt.Sprintf("./credentials/%s.json",answer))
			if err != nil {
				log.Fatal(err)
			}
		},
	}
}