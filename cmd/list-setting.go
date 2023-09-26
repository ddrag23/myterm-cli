package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"termius-cli/entity"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

func checkboxes(label string, opts []string) string {
	var res string
	prompt := &survey.Select{
		Message: label,
		Options: opts,
	}
	survey.AskOne(prompt, &res)

	return res
}

func ListSetting() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all your credential ssh",
		Run: func(cmd *cobra.Command, args []string) {
			files, err := ioutil.ReadDir("./credentials")
			if err != nil {
				log.Fatal(err)
			}
			var listFiles []string
			for _, file := range files {
				jsonFile, err := os.Open("./credentials/" + file.Name())
				if err != nil {
					fmt.Println(err.Error())
				}
				defer jsonFile.Close()
				var cred entity.SSHCredentials
				byteValue, _ := ioutil.ReadAll(jsonFile)
				json.Unmarshal(byteValue, &cred)
				listFiles = append(listFiles, cred.Label)
			}
			answer := checkboxes("Pilih salah satu credential", listFiles)
			RunAndConnectSsh("./credentials/" + answer + ".json")
		},
	}
}
