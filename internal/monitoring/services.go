package monitoring

import (
	"LinMon/internal/ssh_con"
	"bytes"
	"fmt"
	"strings"
	"text/tabwriter"

	"golang.org/x/crypto/ssh"
)

func ServicesUpdater(client *ssh.Client) string {
	out, err := ssh_con.RunCommand(client, "systemctl list-units --type=service --state=running,failed --no-legend --no-pager | awk '{print $4, $1, $2}'\n")
	if err != nil {
		return fmt.Sprintf("Error: %v", err)
	}

	lines := strings.Split(out, "\n")

	var buf bytes.Buffer

	writer := tabwriter.NewWriter(
		&buf,
		0, 0, 1, ' ', tabwriter.Debug,
	)

	// نوشتن عنوان ستون‌ها
	fmt.Fprintln(writer, "STATUS\tSERVICE\tSUB")

	for _, line := range lines {
		if len(line) > 0 {
			fields := strings.Fields(line)
			if len(fields) == 3 {
				status := fields[0]
				service := fields[1]
				sub := fields[2]
				fmt.Fprintf(writer, "%s\t%s\t%s\n", status, service, sub)
			}
		}
	}

	writer.Flush()

	return buf.String()
}
