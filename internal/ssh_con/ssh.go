package ssh_con

import (
	"LinMon/internal/config"
	"flag"
	"fmt"
	"github.com/kbolino/pageant"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
	"log"
	"strings"
)

// parseSSHFlag parses the command-line flags for SSH connection details and returns the username, server, and password.
func ParseSSHFlag() (*string, *string, *string, error) {
	// Define command-line flags for SSH connection and password
	sshFlag := flag.String("ssh", "", "Specify the SSH connection in the format username@server")
	passwordFlag := flag.String("password", "", "Specify the password for the SSH connection")

	// Parse the command-line flags
	flag.Parse()

	// Check if the SSH flag is provided
	if *sshFlag == "" {
		return nil, nil, nil, fmt.Errorf("Error: No SSH connection specified")
	}

	// Split the SSH flag into username and server parts
	sshParts := strings.Split(*sshFlag, "@")
	if len(sshParts) != 2 {
		return nil, nil, nil, fmt.Errorf("Error: Invalid SSH connection format")
	}

	username := sshParts[0]
	server := sshParts[1]

	return &username, &server, passwordFlag, nil
}

// parseSSHConfigYml is a placeholder function for parsing SSH configuration from a YAML file.
func ParseSSHConfigYml() (*config.Config, error) {
	configs, err := config.YamlLoad()
	if err != nil {
		return nil, fmt.Errorf("Not implemented")
	}

	return configs, nil
}

// connectSSH establishes an SSH connection to the specified server using the provided username and password.
func ConnectSSH(username string, server string, password *string) (*ssh.Client, error) {
	var signers []ssh.Signer
	var auth []ssh.AuthMethod

	// Determine the authentication method based on the presence of a password
	if password == nil || *password == "" {
		// Use SSH agent if no password is provided
		agentConn, err := pageant.NewConn()
		if err != nil {
			log.Fatalf("Error connecting to pageant: %v", err)
		}

		sshAgent := agent.NewClient(agentConn)
		signers, err = sshAgent.Signers()
		if err != nil {
			log.Fatalf("Error loading private key: %v", err)
		}

		auth = []ssh.AuthMethod{
			ssh.PublicKeys(signers...),
		}
	} else {
		// Use password authentication if a password is provided
		auth = []ssh.AuthMethod{
			ssh.Password(*password),
		}
	}

	// Configure the SSH client
	conf := &ssh.ClientConfig{
		User:            username,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth:            auth,
	}

	// Dial the SSH server
	client, errSSH := ssh.Dial("tcp", server+":22", conf)
	if errSSH != nil {
		return nil, errSSH
	}

	return client, nil
}

// RunCommand executes a command on the remote SSH server and returns the output or any error encountered.
func RunCommand(client *ssh.Client, command string) (string, error) {
	session, err := client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close() // Ensure the session is closed after execution

	// Execute the command and capture the output
	output, err := session.CombinedOutput(command)
	if err != nil {
		return "", err
	}
	return string(output), nil
}
