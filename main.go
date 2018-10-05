package main

import (
	"os"

	"github.com/berfinsari/tracebeat/cmd"

	_ "github.com/berfinsari/tracebeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
