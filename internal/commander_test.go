package internal_test

import (
	"io"
	"log/slog"
	"testing"

	"github.com/linkinlog/jobbr/internal"
)

func TestCommander_Ping(t *testing.T) {
	c := internal.NewCommander(slog.New(slog.NewJSONHandler(io.Discard, nil)))

	result, err := c.Ping("localhost")

	if err != nil {
		t.Fatal("Ping failed", err)
	}
	if !result.Successful {
		t.Fatal("Ping unsuccessful")
	}
}

func TestCommander_GetSystemInfo(t *testing.T) {
	c := internal.NewCommander(slog.New(slog.NewJSONHandler(io.Discard, nil)))

	info, err := c.GetSystemInfo()

	if err != nil {
		t.Fatal("Failed to get system info", err)
	}
	if info.Hostname == "" {
		t.Fatal("Hostname is empty")
	}
	if info.IPAddress == "" {
		t.Fatal("IP Address is empty")
	}
}
