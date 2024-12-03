package main

import (
	"context"
	"log/slog"
	"os"
)

type runnable func(ctx context.Context) int

var commands = map[string]runnable{
	"serve": serve,
	"setup": setup,
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		slog.Error("A command is required")
		os.Exit(1)
	}

	cmd := args[0]
	runner, ok := commands[cmd]
	if !ok {
		slog.Error("Unknown command", "command", cmd)
	}

	if os.Getenv("DEBUG") != "" {
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}

	os.Exit(runner(context.Background()))
}
