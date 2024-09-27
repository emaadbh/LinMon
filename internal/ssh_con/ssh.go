package ssh_con

import (
	"flag"
	"fmt"
	"github.com/kbolino/pageant"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
	"log"
	"strings"
)

func Ssh() (string, *ssh.Client, error) {
	user, server, password, err := parseSSHFlag()
	if err != nil {
		return *server, nil, err
	}

	return connectSSH(*user, *server, password)
}

func parseSSHFlag() (*string, *string, *string, error) {
	sshFlag := flag.String("ssh", "", "Specify the SSH connection in the format username@server")
	passwordFlag := flag.String("password", "", "Specify the password for the SSH connection")

	flag.Parse()

	if *sshFlag == "" {
		return nil, nil, nil, fmt.Errorf("Error: No SSH connection specified")
	}

	sshParts := strings.Split(*sshFlag, "@")

	if len(sshParts) != 2 {
		return nil, nil, nil, fmt.Errorf("Error: No SSH connection specified")
	}

	username := sshParts[0]
	server := sshParts[1]

	return &username, &server, passwordFlag, nil

}

func parseSSHConfigYml() (*string, *string, *string, error) {

	return nil, nil, nil, fmt.Errorf("Not implemented")
}

func connectSSH(username string, server string, password *string) (string, *ssh.Client, error) {
	var signers []ssh.Signer
	var auth []ssh.AuthMethod

	if password == nil || *password == "" {

		agentConn, err := pageant.NewConn()

		sshAgent := agent.NewClient(agentConn)
		signers, err = sshAgent.Signers()

		if err != nil {
			log.Fatalf("Error loading private key: %v", err)
		}

		auth = []ssh.AuthMethod{
			//ssh.Password(password),
			ssh.PublicKeys(signers...),
		}
	} else {
		auth = []ssh.AuthMethod{
			ssh.Password(*password),
		}
	}

	conf := &ssh.ClientConfig{
		User:            username,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth:            auth,
	}

	client, errSSH := ssh.Dial("tcp", server+":22", conf)
	if errSSH != nil {
		return server, nil, errSSH
	}

	return server, client, nil
}

func RunCommand(client *ssh.Client, command string) (string, error) {
	session, err := client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()

	output, err := session.CombinedOutput(command)
	if err != nil {
		return "", err
	}
	return string(output), nil
}
