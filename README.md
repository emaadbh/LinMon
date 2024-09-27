# Linux Monitoring (LinMon)

**Linux Monitoring (LinMon)** is a lightweight, command-line tool written in Go, designed to monitor key system resources such as CPU usage, memory consumption, disk usage, and running processes on Linux servers.

## Features

- **CPU Usage Monitoring**: Uses the `mpstat` command to fetch CPU usage statistics.
- **Memory Monitoring**: Gathers memory usage information using the `free` command.
- **Disk Usage Monitoring**: Displays disk usage via the `df` command.
- **Process Monitoring**: Retrieves a list of running processes using the `ps` command.
- **Timed Monitoring**: Continuously monitors the system at regular intervals (every 10 seconds by default).

## Usage

Run the executable to start monitoring your server:

```bash
./server-monitor
