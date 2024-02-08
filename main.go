package main

import (
	"os"

	"github.com/tokend/sandwich-bot/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
