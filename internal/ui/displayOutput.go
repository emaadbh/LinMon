package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"time"
)

func DisplayOutput(app *tview.Application, title string, updaterFunc func() string) *tview.TextView {

	outputBox := tview.NewTextView()

	outputBox.
		SetTitle(title).
		SetTitleAlign(tview.AlignLeft).
		SetTitleColor(tcell.ColorGreen).
		SetBackgroundColor(tcell.Color17).
		SetBorder(true)

	go func() {
		for {
			out := updaterFunc()
			outputBox.SetText(out)

			app.Draw()
			time.Sleep(1 * time.Second)
		}

	}()

	outputBox.SetBorder(true)

	return outputBox
}
