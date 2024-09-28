package main

import (
	"LinMon/internal/monitoring"
	"LinMon/internal/ssh_con"
	"LinMon/internal/ui"
	"github.com/rivo/tview"
	"golang.org/x/crypto/ssh"
	"log"
)

func main() {
	// Connect to SSH server
	_, client, err := ssh_con.Ssh()
	if err != nil {
		log.Fatalf("SSH connection error: %v", err)
		return
	}
	app := tview.NewApplication()

	outputBoxes := createOutputBoxes(client, app)

	ui.InitializeUI(app, outputBoxes)
}

func createOutputBoxes(client *ssh.Client, app *tview.Application) []tview.Primitive {
	return []tview.Primitive{
		createOutputBox(app, "Network", monitoring.NetworkUpdater, client),
		createOutputBox(app, "TOP", monitoring.CpuUpdater, client),
		createOutputBox(app, "Journal LOG", monitoring.JournalUpdater, client),
		createOutputBox(app, "WebServer", monitoring.WebServerUpdater, client),
	}
}

func createOutputBox(app *tview.Application, label string, updater func(*ssh.Client) string, client *ssh.Client) tview.Primitive {
	return ui.DisplayOutput(app, label, func() string {
		return updater(client)
	})
}
