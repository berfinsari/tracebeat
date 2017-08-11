package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"

	"github.com/berfinsari/tracebeat/beater"
)

func main() {
	err := beat.Run("tracebeat", "", beater.New)
	if err != nil {
		os.Exit(1)
	}
}
