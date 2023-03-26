package cli

import (
	"github.com/DataDrake/cli-ng/v2/cmd"
	log "github.com/DataDrake/waterlog"
	"github.com/DataDrake/waterlog/format"
	"github.com/DataDrake/waterlog/level"
)

func init() {
	log.SetFormat(format.Min)
	log.SetLevel(level.Debug)
	cmd.Register(&cmd.Help)
	cmd.Register(&cmd.Version)
}

// Root command
var Root = cmd.Root{
	Name:      "repo-abi-scanner",
	Short:     "Tool for quickly generating ABI reports for source repos",
	Version:   "0.5.0",
	Copyright: "© 2021 Solus Project <copyright@getsol.us>",
}
