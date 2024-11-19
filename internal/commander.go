package internal

import (
	"log/slog"
	"net"
	"os"
	"os/exec"
	"time"
)

type commander struct {
	logger *slog.Logger
}

func NewCommander(l *slog.Logger) Commander {
	return &commander{
		logger: l,
	}
}

func (c *commander) Ping(host string) (PingResult, error) {
	c.logger.Debug("Pinging host", "host", host)

	cmd := exec.Command("ping", "-c", "1", host)
	start := time.Now()
	err := cmd.Run()

	c.logger.Info("Ping completed", "host", host, "duration", time.Since(start), "success", err == nil)

	return PingResult{
		Successful: err == nil,
		Time:       time.Since(start),
	}, err
}

func (c *commander) GetSystemInfo() (SystemInfo, error) {
	c.logger.Debug("Getting system info")

	hostname, err := os.Hostname()
	if err != nil {
		return SystemInfo{}, err
	}

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return SystemInfo{}, err
	}

	var ip string
	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ip = ipNet.IP.String()
				break
			}
		}
	}

	c.logger.Info("Got system info", "hostname", hostname, "ip", ip)

	return SystemInfo{Hostname: hostname, IPAddress: ip}, nil
}
