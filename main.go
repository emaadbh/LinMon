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
	var host *string
	var user *string
	var password *string

	user, host, password = getDataHosts()

	client, err := ssh_con.ConnectSSH(*user, *host, password)
	if err != nil {
		log.Fatalf("SSH connection error: %v", err)
		return
	}
	app := tview.NewApplication()

	outputBoxes := createOutputBoxes(client, app)

	ui.InitializeUI(*host, app, outputBoxes)
}

func getDataHosts() (user *string, host *string, password *string) {

	user, host, password, err := ssh_con.ParseSSHFlag()

	if err != nil {
		configs, errYaml := ssh_con.ParseSSHConfigYml()
		if errYaml != nil {
			log.Fatalf("SSH connection error: %v", err)
			return
		}

		for _, config := range configs.Servers {
			host = &config.IP
			user = &config.User
			password = &config.Password
			break
		}
	}

	return user, host, password
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
