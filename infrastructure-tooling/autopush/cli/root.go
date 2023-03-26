package cli

import (
	"github.com/DataDrake/cli-ng/cmd"
)

var Root = cmd.Root{
	Name:   "autopush",
	Short:  "tool for automatically pushing stack rebuilds",
	Flags:  &RootFlags{},
	Single: false,
}

type RootFlags struct{}
