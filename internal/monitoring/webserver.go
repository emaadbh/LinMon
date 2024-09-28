package monitoring

import (
	"LinMon/internal/ssh_con"
	"fmt"
	"golang.org/x/crypto/ssh"
	"strings"
)

func WebServerUpdater(client *ssh.Client) string {
	webCMD := webserver(client)

	out, err := ssh_con.RunCommand(client, webCMD)
	if err != nil {
		out = fmt.Sprintf("Error: %v", err)
	}

	return out
}

// todo: need improvement
func webserver(client *ssh.Client) string {
	cmd := "journalctl -u nginx -n 20 -e"

	webserverStatus := checkWebServer(client)
	if webserverStatus == -1 {
		cmd = "echo 'WEBSERVER: OFF'"

	} else if webserverStatus == 0 {
		cmd = "journalctl -u apache2 -n 20 -e"
	}

	return cmd
}

func checkWebServer(client *ssh.Client) int {

	out, _ := ssh_con.RunCommand(client, "ss -tlnp | grep -oP '(?<=users:\\(\\(\")\\w+(?=\")' | sort -u\n")

	if strings.Contains(out, "apache") {
		return 0
	} else if strings.Contains(out, "nginx") {
		return 1
	}

	return -1
}
