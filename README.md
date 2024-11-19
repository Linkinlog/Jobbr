# Jobbr

**Jobbr** is a simple cross-platform system utility designed to:
- Execute system commands (e.g., ping, system info).
- Run as a daemon on macOS for always-on availability.

This repository includes a macOS installer (`.pkg`) that installs the Jobbr binary to `/usr/local/bin` and configures a Launch Daemon to start Jobbr on boot.

## Features
- **Ping Support**: Send a network ping to a target host.
- **System Info**: Retrieve the hostname and IP address of the system.
- **HTTP API**: Simple JSON-based API for remote command execution.

## Installation

### macOS
1. **Download and Install:**
   - Download the `jobbr.pkg` file from the [Releases](https://github.com/linkinlog/jobbr/releases).
   - Double-click to install.

2. **What the Installer Does:**
   - Installs the `jobbr` binary to `/usr/local/bin/`.
   - Copies the `jobbr.plist` Launch Daemon file to `/Library/LaunchDaemons/`.

3. **Start on Boot:**
   - The installer enables Jobbr to start on system boot using the Launch Daemon.

### Verification
After installation, logs will be generated in `/tmp/jobbr.log`.

You can verify the installation by checking the Jobbr binary and Launch Daemon file:
```bash
which jobbr
# Output: /usr/local/bin/jobbr

ls /Library/LaunchDaemons/com.github.linkinlog.jobbr.plist
# Output: /Library/LaunchDaemons/com.github.linkinlog.jobbr.plist
```

## Usage
![Jobbr Usage](./docs/jobbr.gif)
### Starting Jobbr
Jobbr is automatically started on boot via the Launch Daemon. To manually restart it:

```bash
sudo launchctl unload /Library/LaunchDaemons/com.github.linkinlog.jobbr.plist
sudo launchctl load /Library/LaunchDaemons/com.github.linkinlog.jobbr.plist
```

### HTTP API
Jobbr exposes a simple HTTP API for executing commands.

Endpoint: /execute
- Method: POST

- Request Body:

```json
{
    "type": "ping", // or "sysinfo"
    "payload": "example.com" // For "ping", this is the host; "sysinfo" has no payload
}
```
- Response:

```json
{
    "success": true,
    "data": {...},
    "error": ""
}
```
Example
- To get system info:

```bash
curl -X POST -H "Content-Type: application/json" -d '{"type":"sysinfo"}' http://localhost:59152/execute
```
- Response:

```json
{
    "success": true,
    "data": {
        "hostname": "MyMac",
        "ipaddress": "192.168.1.10"
    },
    "error": ""
}
```
## Building Jobbr
### Prerequisites
- Go installed (version 1.23.1 or later).
- macOS system with pkgbuild available for creating .pkg installers.

### Steps to Build
1. Clone the Repository:

```bash
git clone https://github.com/linkinlog/jobbr.git
cd jobbr
```
2. Build the Installer:

```bash
make build
```

## Testing
- Run unit tests for Jobbr:
```bash
go test ./...
```
## Contributing
- Fork the repository.
- Create a new branch for your feature.
- Submit a pull request with a description of the changes.

## Enjoy using Jobbr! 🎉
