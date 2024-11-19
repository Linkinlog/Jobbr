package internal

import "time"

type PingResult struct {
	Successful bool
	Time       time.Duration
}

type SystemInfo struct {
	Hostname  string
	IPAddress string
}

type CommandRequest struct {
	Type    string `json:"type"`    // "ping" or "sysinfo"
	Payload string `json:"payload"` // For ping, this is the host
}

type CommandResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error,omitempty"`
}
