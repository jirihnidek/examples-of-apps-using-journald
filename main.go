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
	var opts slogjournal.Options
	opts.Level = slog.LevelDebug
	handler, err := slogjournal.NewHandler(&opts)
	if err != nil {
		panic(err)
	}
	logger := slog.New(handler)
	s := "gopher"
	fmt.Printf("Hello and welcome, %s\n", s)
	logger.Info(fmt.Sprintf("Hello and welcome from info level, %s!\n", s))
	logger.Debug(fmt.Sprintf("Hello and welcome from debug level, %s!\n", s))
}
