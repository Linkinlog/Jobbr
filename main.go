package main

import (
	"flag"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/linkinlog/jobbr/internal"
)

func main() {
	logFile, err := os.OpenFile("/tmp/jobbr.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	slogOpts := &slog.HandlerOptions{}

	verbosePtr := flag.Bool("v", false, "verbose")
	flag.Parse()

	if *verbosePtr {
		slogOpts.Level = slog.LevelDebug
	}

	logger := slog.New(slog.NewJSONHandler(logFile, slogOpts))

	logger.Info("Jobbr Started", "verbose", *verbosePtr)

	c := internal.NewCommander(logger)
	h := internal.NewHandler(c, logger)

	if err := http.ListenAndServe(":59152", h); err != nil {
		logger.Error("failed to listen and serve", "error", err.Error())
	}
}
