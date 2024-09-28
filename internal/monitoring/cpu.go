package monitoring

import (
	"LinMon/internal/ssh_con"
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"strings"
)

func CpuUpdater(client *ssh.Client) string {
	cpuOut := runCommand(client, "echo -n 'CPU cores: ' && grep -c ^processor /proc/cpuinfo")
	memOut := runCommand(client, "echo -n ', Used Memory (RAM): ' && free -m | awk '/Mem:/ {print $3 \"MB of \" $2 \"MB used\"}'")
	psOut := runCommand(client, "ps -eo pid,comm,%cpu,%mem --sort=-%cpu | head -n 20")
	loadOut := runCommand(client, "echo -n 'Load Average: ' && cat /proc/loadavg | awk '{print $1\", \"$2\", \"$3}'")

	result := strings.TrimSpace(cpuOut) + strings.TrimSpace(memOut) + "\n" + strings.TrimSpace(loadOut) + "\n" + strings.TrimSpace(psOut)
	return result
}

func runCommand(client *ssh.Client, cmd string) string {
	output, err := ssh_con.RunCommand(client, cmd)
	if err != nil {
		log.Printf("Error running command '%s': %v", cmd, err)
		return fmt.Sprintf("Error: %v", err)
	}
	return output
}
