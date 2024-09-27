package monitoring

import (
	"LinMon/internal/ssh_con"
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"golang.org/x/crypto/ssh"
	"time"
)

func DisplayOutput(client *ssh.Client, app *tview.Application, outputBox *tview.TextView) {
	outputBox.
		SetTitle(" TOP ").
		SetTitleAlign(tview.AlignLeft).
		SetTitleColor(tcell.ColorGreen).
		SetBackgroundColor(tcell.Color17).
		SetBorder(true)

	go func() {
		for {
			out, err := ssh_con.RunCommand(client, "top -b -n 1 | head -n 20")

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
