package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"log"
)

func InitializeUI(app *tview.Application, outBoxes []tview.Primitive) {
	focusedIndex := 0
	mainFlex, focusableItems := setupUI(app, outBoxes)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyF10:
			app.Stop()
		case tcell.KeyTab:
			focusedIndex = (focusedIndex + 1) % len(focusableItems)
			app.SetFocus(focusableItems[focusedIndex])
		}
		return event
	})

	if err := app.SetRoot(mainFlex, true).Run(); err != nil {
		log.Fatalf("Error starting application: %v", err)
	}
}

// setupUI initializes the user interface components and returns the main layout and focusable items
func setupUI(app *tview.Application, outBoxes []tview.Primitive) (*tview.Frame, []tview.Primitive) {
	outputBoxes := outBoxes
	list := createOptionsList(app)

	focusableItems := append(outputBoxes, list)
	flex := createMainFlex(outputBoxes, list)
	mainFlex := createFrame(flex)

	return mainFlex, focusableItems
}

// createFrame creates a new frame for the application UI with server information
func createFrame(mainFlex *tview.Flex) *tview.Frame {
	return tview.NewFrame(mainFlex).
		AddText("F10: Exit   F1: LIST VPS   TAB: Change focus", true, tview.AlignLeft, tcell.ColorRed)
}

// createOptionsList creates a list of options for user interaction
func createOptionsList(app *tview.Application) *tview.List {
	list := tview.NewList()
	list.AddItem("List servers", "", 'a', nil).
		AddItem("MODE Pro", "Not implemented :)", 'b', nil)
	return list
}

// createMainFlex creates the main layout of the UI using the output boxes and options list
func createMainFlex(outputBoxes []tview.Primitive, list *tview.List) *tview.Flex {
	rowFlex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(outputBoxes[0], 0, 2, false).
		AddItem(outputBoxes[1], 0, 2, false).
		AddItem(list, 5, 1, true)

	row2Flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(outputBoxes[2], 0, 1, false).
		AddItem(outputBoxes[3], 0, 1, false)

	return tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(rowFlex, 0, 1, true).
		AddItem(row2Flex, 0, 2, false)
}
