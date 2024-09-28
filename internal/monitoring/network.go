package monitoring

import (
	"LinMon/internal/ssh_con"
	"fmt"
	"golang.org/x/crypto/ssh"
)

func NetworkUpdater(client *ssh.Client) string {
	out, err := ssh_con.RunCommand(client, "ss -plunt")
	if err != nil {
		out = fmt.Sprintf("Error: %v", err)
	}

	return out
}
