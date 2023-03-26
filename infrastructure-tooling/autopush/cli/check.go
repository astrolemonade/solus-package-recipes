package cli

import (
	"dev.getsol.us/sources/infrastructure-tooling/autopush/eopkg"
	"dev.getsol.us/sources/infrastructure-tooling/autopush/order"
	"dev.getsol.us/sources/infrastructure-tooling/autopush/ypkg"
	"github.com/DataDrake/cli-ng/cmd"
	log "github.com/DataDrake/waterlog"
	"github.com/DataDrake/waterlog/format"
	"github.com/DataDrake/waterlog/level"
)

var Check = cmd.Sub{
	Name:  "check",
	Alias: "ck",
	Short: "Verify that all builds have been pushed",
	Flags: &CheckFlags{},
	Args:  &CheckArgs{},
	Run:   CheckRun,
}

func init() {
	cmd.Register(&Check)
}

type CheckFlags struct{}

type CheckArgs struct {
	ConfigFile string `desc:"path to Build Order config file"`
}

func CheckRun(r *cmd.Root, c *cmd.Sub) {
	args := c.Args.(*CheckArgs)

    log.SetFormat(format.Min)
    log.SetLevel(level.Info)

	f, err := order.Load(args.ConfigFile)
	if err != nil {
		log.Fatalf("Failed to parse config file '%s', reason '%s'\n", args.ConfigFile, err)
	}

	log.Infoln("Fetching index...")
	i, err := eopkg.FetchIndex()
	if err != nil {
		log.Fatalf("Failed to fetch index: '%s'\n", err)
	}
	log.Goodln("Index updated.")
    for {
		t, ok := f.Next()
		if !ok {
			break
		}
		log.Infof("[%s]\n", t.Name)
		for {
			dir, ok := t.Next()
			if !ok {
				break
			}
			log.Infof("\tPackage: %s\n", dir)
			log.Infoln("\t\tLoading package.yml...")
			pkg, err := ypkg.Load(dir)
			if err != nil {
				log.Fatalf("Faied to read 'package.yml', reason: %s\n", err)
			}
			log.Goodln("\t\tOK.")
			log.Infof("\t\tFound '%s', release '%d'.\n", pkg.Name, pkg.Release)

			if i.HasRelease(pkg.Name, pkg.Release) {
				log.Goodf("\t\tFound in Index.\n\n")
				continue
			}
			log.Errorf("\t\tRelease '%d' not found\n\n", pkg.Release)
		}
		log.Goodf("[Done]\n\n")
	}
}
