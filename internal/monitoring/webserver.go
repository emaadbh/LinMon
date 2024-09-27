package monitoring

import (
	"LinMon/internal/ssh_con"
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"golang.org/x/crypto/ssh"
	"strings"
	"time"
)

func WebServerOutput(client *ssh.Client, app *tview.Application, outputBox *tview.TextView) {
	webserver := checkWebServer(client)
	cmd := "journalctl -u nginx -n 20 -e"

	if webserver == -1 {
		outputBox.SetTitle("WEBSERVER: OFF")
		cmd = "echo 'WEBSERVER: OFF'"

	} else if webserver == 0 {
		outputBox.SetTitle("WEBSERVER: ON (Apache)")
		cmd = "journalctl -u apache2 -n 20 -e"

	} else if webserver == 1 {
		outputBox.SetTitle("WEBSERVER: ON (Nginx)")
	}
	outputBox.SetTitleAlign(tview.AlignLeft).
		SetTitleColor(tcell.ColorRed).
		SetBackgroundColor(tcell.ColorGold).
		SetBorder(true)

	go func() {
		for {
			out, err := ssh_con.RunCommand(client, cmd)

			if err != nil {
				out = fmt.Sprintf("Error: %v", err)
			}
			outputBox.SetText(out)

			app.Draw()
			time.Sleep(1 * time.Second)
		}

	}()

	outputBox.SetBorder(true)
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
