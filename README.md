# LinMon: Linux Server Monitoring Tool (In Development)

Linux Monitoring (LinMon)** is a lightweight, command-line tool written in Go, designed to monitor key system resources such as CPU usage , This tool is under active development, with additional features planned for future releases.

## Current Features

1. **SSH Connection**: Securely connects to a Linux server via SSH.
2. **Journal Logs Monitoring**: Fetches and displays system logs using `journalctl`.
3. **Network Status Monitoring**: Monitors active network connections and socket states using `ss`.
4. **Resource Usage Monitoring**: Tracks CPU and system resource usage through the `top` command.
5. **Web Server Log Monitoring**: Automatically detects and displays logs for **Apache** or **Nginx** web servers via `journalctl` if installed on the server.

## How to Run

### Method 1: Using Command-Line Parameters
To execute LinMon and connect to a Linux server via SSH, you can use the following command, replacing the SSH connection details with your server's information:


```bash
go run main.go --ssh root@192.168.1.111 --password 12345
```
### Method 2: Using config.yml File
Alternatively, you can use a configuration file to manage your server connection details. This allows you to avoid passing parameters like --ssh and --password directly on the command line. The config.yml file should be located in the configs/ directory and structured as follows:

```yaml
servers:
  vps1_web:
    user: root
    ip: 192.168.1.1
    port: 22
    password: pass123
  vps2_db:
    user: root
    ip: 192.168.1.2
    port: 22
    password: pass456
```


### Key-Based Authentication with Pageant
If Pageant is running on your system and your private SSH key is loaded, the package will automatically use the key for authentication. There is no need to manually input a password in this case.

To use Pageant:
Ensure Pageant is running and your key is loaded.
Simply run the application with the --ssh flag. The package will use the key-based authentication method without needing the --password flag.



## Upcoming Features

1. **Memory Monitoring**: Gather memory usage information using the `free` command.
2. **Disk Usage Monitoring**: Display disk usage statistics using the `df` command.
3. **Process Monitoring**: Retrieve a list of active processes using the `ps` command.
4. **Config File Support**: Define server connection details through a `config.yaml` file.
5. **Multiple Server Management**: Manage and monitor multiple servers from `config.yaml`.
6. **Advanced Log Monitoring**: More detailed and sophisticated log analysis for system and application logs.
7. **Custom SSH Port**: Ability to connect using a specified SSH port.

---

