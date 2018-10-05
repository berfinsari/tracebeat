package cmd

import (
	"github.com/berfinsari/tracebeat/beater"

	cmd "github.com/elastic/beats/libbeat/cmd"
)

// Name of this beat
var Name = "tracebeat"

// RootCmd to handle beats cli
var RootCmd = cmd.GenRootCmd(Name, "", beater.New)
