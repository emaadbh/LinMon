package monitoring

import (
	"LinMon/internal/ssh_con"
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"golang.org/x/crypto/ssh"
	"time"
)

func NetworkDisplay(client *ssh.Client, app *tview.Application, outputBox *tview.TextView) {
	outputBox.
		SetTitle(" NETWORK (SS)").
		SetTitleAlign(tview.AlignLeft).
		SetTitleColor(tcell.ColorBlue).
		SetBackgroundColor(tcell.Color18).
		SetBorder(true)

	go func() {
		for {
			out, err := ssh_con.RunCommand(client, "ss -plunt")

			if err != nil {
				out = fmt.Sprintf("Error: %v", err)
			}
			outputBox.SetText(out)

			app.Draw()
			time.Sleep(3 * time.Second)
		}

	}()

	outputBox.SetBorder(true)
}
