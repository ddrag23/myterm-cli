package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"termius-cli/entity"
	"termius-cli/utils"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

func AddSetting() *cobra.Command {
	return &cobra.Command{
		Use:   "add",
		Short: "setting your credentials ssh",
		Run: func(c *cobra.Command, args []string) {
			var qs = []*survey.Question{
				{
					Name:     "label",
					Prompt:   &survey.Input{Message: "Input label ssh?"},
					Validate: survey.Required,
				},
				{
					Name:     "address",
					Prompt:   &survey.Input{Message: "Input address ssh?"},
					Validate: survey.Required,
				},
				{
					Name:     "port",
					Prompt:   &survey.Input{Message: "Input port ssh?", Default: "22"},
					Validate: survey.Required,
				},
				{
					Name:     "username",
					Prompt:   &survey.Input{Message: "Input username ssh?"},
					Validate: survey.Required,
				},
				{
					Name:     "password",
					Prompt:   &survey.Input{Message: "Input password ssh?"},
					Validate: survey.Required,
				},
			}
			var answers entity.SSHCredentials
			err := survey.Ask(qs, &answers)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			file, _ := json.MarshalIndent(answers, "", " ")
			checkDir, err := utils.Exists("./credentials")
			if err != nil {
				log.Fatal(err.Error())
			}
			if !checkDir {
				os.MkdirAll("credentials", 0777)
			}
			existFile, err := utils.Exists(fmt.Sprintf("./credentials/%s.json", answers.Label))
			if err != nil {
				log.Fatal(err)
			}
			if existFile {
				log.Fatalln("your label input is exist")
			}
			err = ioutil.WriteFile("./credentials/"+answers.Label+".json", file, 0777)
			if err != nil {
				fmt.Println(err.Error())
			}
		},
	}
}
