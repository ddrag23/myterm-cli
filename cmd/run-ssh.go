package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"termius-cli/entity"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

func runAndConnectSsh(file string) {
	jsonFile, _ := os.Open(file)
	defer jsonFile.Close()
	var cred entity.SSHCredentials
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &cred)
	port := strconv.Itoa(cred.Port)

	sshConfig := &ssh.ClientConfig{
		User:            cred.Username,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth: []ssh.AuthMethod{
			ssh.Password(cred.Password),
		},
	}
	client, err := ssh.Dial("tcp", cred.Address+":"+port, sshConfig)
	if client != nil {
		defer client.Close()
	}
	if err != nil {
		log.Fatal("Failed to dial. " + err.Error())
	}
	session, err := client.NewSession()
	if session != nil {
		defer session.Close()
	}
	if err != nil {
		log.Fatal("Failed to create session. " + err.Error())
	}
	session.Stdin = os.Stdin
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	modes := ssh.TerminalModes{
		ssh.ECHO:          1,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}
	fileDescriptor := int(os.Stdin.Fd())

	if terminal.IsTerminal(fileDescriptor) {
		originalState, err := terminal.MakeRaw(fileDescriptor)
		if err != nil {
			log.Fatal(err)
		}
		defer terminal.Restore(fileDescriptor, originalState)

		termWidth, termHeight, err := terminal.GetSize(fileDescriptor)
		if err != nil {
			log.Fatal(err)

		}

		err = session.RequestPty("xterm-256color", termHeight, termWidth, modes)
		if err != nil {
			log.Fatal(err)

		}
	}
	// err = session.RequestPty("xterm-256color", 80, 100, modes)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = session.Shell()
	if err != nil {
		log.Fatal(err)
	}
	session.Wait()
	defer session.Close()
}

func RunSsh() *cobra.Command {
	return &cobra.Command{
		Use:   "run",
		Short: "run ssh",
		Run: func(c *cobra.Command, args []string) {
			runAndConnectSsh(fmt.Sprintf("./credentials/%s.json", args[0]))
		},
	}
}
