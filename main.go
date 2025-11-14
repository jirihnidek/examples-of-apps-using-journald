package main

import (
	"fmt"
	"log/slog"

	slogjournal "github.com/systemd/slog-journal"
)

// Example of the client application using slogjournal writing logs to systemd journal
//
// More documentation can be found at:
// https://pkg.go.dev/github.com/systemd/slog-journal

func main() {
	// Configure journald handler
	var opts slogjournal.Options
	opts.Level = slog.LevelDebug
	handler, err := slogjournal.NewHandler(&opts)
	if err != nil {
		panic(err)
	}

	// Create the new logger using journald for logging
	logger := slog.New(handler)

	// Use the logger for some logging
	s := "gopher"
	logger.Info(fmt.Sprintf("Hello from info level, %s!\n", s))
	logger.Debug(fmt.Sprintf("Hello from debug level, %s!\n", s))

	// This is an example of using custom keys.
	// The KEY will not be displayed in the e.g. short output mode, but
	// it will be displayed in the e.g. verbose or json* mode.
	// The journal only accepts keys of the form ^[A-Z_][A-Z0-9_]*$.
	logger.Debug(
		"Hello from debug level with custom KEY",
		slog.String("KEY", "value"),
	)
}
