package cli

import (
	"github.com/DataDrake/cli-ng/v2/cmd"
	log "github.com/DataDrake/waterlog"
	"repo-abi-scanner/pkg"
)

func init() {
	cmd.Register(Single)
}

// Single generates the ABI reports for a single repo
var Single = &cmd.Sub{
	Name:  "single",
	Short: "Generate ABI reports for a single package",
	Args:  &SingleArgs{},
	Run:   SingleRun,
}

// SingleArgs are the argument(s) for the "single" sub-command
type SingleArgs struct {
	Source string `desc:"Path of source Git repo for generating reports"`
	Repo   string `desc:"Path of package repo for generating reports"`
}

// SingleRun carries out the "single" sub-command
func SingleRun(r *cmd.Root, s *cmd.Sub) {
	args := s.Args.(*SingleArgs)
	src, err := pkg.NewSource(args.Source)
	if err != nil {
		log.Fatalf("Failed to load Source, reason: %s\n", err)
	}
	if err = src.Scan(args.Repo); err != nil {
		log.Fatalf("Failed to scan Source, reason: %s\n", err)
	}
	if err = src.Save(); err != nil {
		log.Fatalf("Failed to save ABI reports, reason: %s\n", err)
	}
	if err = src.Commit(); err != nil {
		log.Fatalf("Failed to commit ABI reports, reason: %s\n", err)
	}
	log.Goodln("Done")
}
