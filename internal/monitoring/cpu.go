package monitoring

import (
	"LinMon/internal/ssh_con"
	"fmt"
	"golang.org/x/crypto/ssh"
)

func CpuUpdater(client *ssh.Client) string {
	out, err := ssh_con.RunCommand(client, "top -b -n 1 | head -n 20")
	if err != nil {
		out = fmt.Sprintf("Error: %v", err)
	}

	return out
}
