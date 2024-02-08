package main

import (
	"os"

	"sandwich-bot/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
