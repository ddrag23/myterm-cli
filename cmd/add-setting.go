package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"termius-cli/entity"
	"termius-cli/utils"

	"github.com/spf13/cobra"
)

func AddSetting() *cobra.Command {
	return &cobra.Command{
		Use:   "add-setting",
		Short: "setting your credentials ssh",
		Run: func(c *cobra.Command, args []string) {
			label := utils.StringPrompt("Input label ssh ")
			address := utils.StringPrompt("Input address ssh ")
			port := utils.StringPrompt("Input port ssh, default using port 22")
			username := utils.StringPrompt("Input username ssh")
			password := utils.StringPrompt("Input password ssh")
			var convertPort int
			if port == "" {
				convertPort = 22
			} else {
				convertPort, _ = strconv.Atoi(port)
			}
			input := entity.SSHCredentials{
				Label:    label,
				Address:  address,
				Port:     convertPort,
				Username: username,
				Password: password,
			}
			file, _ := json.MarshalIndent(input, "", " ")
			checkDir, err := utils.Exists("./credentials")
			if err != nil {
				log.Fatal(err.Error())
			}
			if !checkDir {
				os.MkdirAll("credentials", 0777)
			}
			err = ioutil.WriteFile("./credentials/"+label+".json", file, 0777)
			if err != nil {
				fmt.Println(err.Error())
			}
		},
	}
}
