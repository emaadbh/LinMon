package main

import (
	"log"

	"LinMon/internal/monitoring"
	"LinMon/internal/ssh_con"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"golang.org/x/crypto/ssh"
)

func main() {
	focusedIndex := 0

	// Connect to SSH server
	server, client, err := ssh_con.Ssh()
	if err != nil {
		log.Fatalf("SSH connection error: %v", err)
		return
	}

	// Create a new TView application
	app := tview.NewApplication()

	// Setup the user interface and get the main layout and focusable items
	mainFlex, focusableItems := setupUI(client, app, server)

	// Capture keyboard input for application control
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyF10:
			app.Stop()
		case tcell.KeyCtrlA:
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
func setupUI(client *ssh.Client, app *tview.Application, server string) (*tview.Frame, []tview.Primitive) {
	outputBoxes := createOutputBoxes(client, app)
	list := createOptionsList(app)

	focusableItems := append(outputBoxes, list)
	flex := createMainFlex(outputBoxes, list)
	mainFlex := createFrame(server, flex)

	return mainFlex, focusableItems
}

// createFrame creates a new frame for the application UI with server information
func createFrame(server string, mainFlex *tview.Flex) *tview.Frame {
	return tview.NewFrame(mainFlex).
		AddText("IP: "+server, true, tview.AlignLeft, tcell.ColorRed).
		AddText("F10: Exit - F1: LIST VPS - CTRL+A: Change focus", true, tview.AlignLeft, tcell.ColorRed)
}

// createOutputBoxes creates text views for different monitoring outputs and updates them
func createOutputBoxes(client *ssh.Client, app *tview.Application) []tview.Primitive {
	outputBoxCpu := tview.NewTextView()
	outputBoxNetwork := tview.NewTextView()
	outputBoxJournal := tview.NewTextView()
	outputBoxWebserver := tview.NewTextView()

	// Call the monitoring functions to update the text views
	monitoring.DisplayOutput(client, app, outputBoxCpu)
	monitoring.NetworkDisplay(client, app, outputBoxNetwork)
	monitoring.JournalDisplay(client, app, outputBoxJournal)
	monitoring.WebServerOutput(client, app, outputBoxWebserver)

	return []tview.Primitive{outputBoxNetwork, outputBoxCpu, outputBoxJournal, outputBoxWebserver}
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
