package monitoring

import (
	"LinMon/internal/ssh_con"
	"LinMon/utils"
	"fmt"
	"golang.org/x/crypto/ssh"
	"strings"
)

func JournalUpdater(client *ssh.Client) string {
	out, err := ssh_con.RunCommand(client, "journalctl -o short -p 0..4 | head -n 20")
	if err != nil {
		out = fmt.Sprintf("Error: %v", err)
	}

	logs := strings.Split(out, "\n")
	output := utils.FilterImportantLog(logs)
	return output
}
