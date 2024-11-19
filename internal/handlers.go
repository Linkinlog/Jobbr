package internal

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type Commander interface {
	Ping(host string) (PingResult, error)
	GetSystemInfo() (SystemInfo, error)
}

func NewHandler(c Commander, l *slog.Logger) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/execute", handleExecute(c, l))
	return mux
}

func handleExecute(c Commander, l *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l.Debug("Received request", "method", r.Method, "url", r.URL.String())

		var req CommandRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			resp := CommandResponse{Error: err.Error()}
			l.Error("Failed to decode request", "error", err, "response", resp)
			if encErr := json.NewEncoder(w).Encode(resp); encErr != nil {
				l.Error("Failed to encode response", "error", encErr)
				http.Error(w, encErr.Error(), http.StatusInternalServerError)
			}
			return
		}

		var resp CommandResponse
		switch req.Type {
		case "ping":
			result, err := c.Ping(req.Payload)
			if err != nil {
				resp = CommandResponse{Error: err.Error()}
			} else {
				resp = CommandResponse{Success: result.Successful, Data: result}
			}
		case "sysinfo":
			info, err := c.GetSystemInfo()
			if err != nil {
				resp = CommandResponse{Error: err.Error()}
			} else {
				resp = CommandResponse{Success: true, Data: info}
			}
		default:
			resp = CommandResponse{Error: "unknown command"}
		}

		l.Info("Sending response", "response", resp)

		if err := json.NewEncoder(w).Encode(resp); err != nil {
			l.Error("Failed to encode response", "error", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
