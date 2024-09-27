# LinMon: Linux Server Monitoring Tool (In Development)

Linux Monitoring (LinMon)** is a lightweight, command-line tool written in Go, designed to monitor key system resources such as CPU usage , This tool is under active development, with additional features planned for future releases.

## Current Features

1. **SSH Connection**: Securely connects to a Linux server via SSH.
2. **Journal Logs Monitoring**: Fetches and displays system logs using `journalctl`.
3. **Network Status Monitoring**: Monitors active network connections and socket states using `ss`.
4. **Resource Usage Monitoring**: Tracks CPU and system resource usage through the `top` command.
5. **Web Server Log Monitoring**: Automatically detects and displays logs for **Apache** or **Nginx** web servers via `journalctl` if installed on the server.

## How to Run

To execute LinMon, use the following command, replacing the SSH connection details with your server's:

```bash
./go run .\main.go --ssh root@192.168.1.111
```


## Upcoming Features

1. **Memory Monitoring**: Gather memory usage information using the `free` command.
2. **Disk Usage Monitoring**: Display disk usage statistics using the `df` command.
3. **Process Monitoring**: Retrieve a list of active processes using the `ps` command.
4. **Config File Support**: Define server connection details through a `config.yaml` file.
5. **Multiple Server Management**: Manage and monitor multiple servers from `config.yaml`.
6. **Advanced Log Monitoring**: More detailed and sophisticated log analysis for system and application logs.
7. **Custom SSH Port**: Ability to connect using a specified SSH port.

---

