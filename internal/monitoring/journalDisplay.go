package monitoring

import (
	"LinMon/internal/ssh_con"
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"golang.org/x/crypto/ssh"
	"time"
)

func JournalDisplay(client *ssh.Client, app *tview.Application, outputBox *tview.TextView) {
	outputBox.
		SetTitle(" Journal LOG ").
		SetTitleAlign(tview.AlignLeft).
		SetTitleColor(tcell.ColorYellow).
		SetBackgroundColor(tcell.Color17).
		SetBorder(true)

	go func() {
		for {
			out, err := ssh_con.RunCommand(client, "journalctl -o short -p 0..4 | head -n 20")

			if err != nil {
				out = fmt.Sprintf("Error: %v", err)
			}
			outputBox.SetText(out)

			app.Draw()
			time.Sleep(5 * time.Second)
		}

	}()

	outputBox.SetBorder(true)
}
