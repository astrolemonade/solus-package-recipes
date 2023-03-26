package cli

import (
	"dev.getsol.us/sources/infrastructure-tooling/autopush/eopkg"
	"dev.getsol.us/sources/infrastructure-tooling/autopush/order"
	"dev.getsol.us/sources/infrastructure-tooling/autopush/ypkg"
	"fmt"
	"github.com/DataDrake/cli-ng/cmd"
	"os"
)

var Dry = cmd.Sub{
	Name:  "dryrun",
	Alias: "dr",
	Short: "Run through a Build Order without actually doing anything",
	Flags: &DryFlags{},
	Args:  &DryArgs{},
	Run:   DryRun,
}

func init() {
	cmd.Register(&Dry)
}

type DryFlags struct{}

type DryArgs struct {
	ConfigFile string `desc:"path to Build Order config file"`
}

func DryRun(r *cmd.Root, c *cmd.Sub) {
	args := c.Args.(*DryArgs)
	f, err := order.Load(args.ConfigFile)
	if err != nil {
		fmt.Printf("Failed to parse config file '%s', reason '%s'\n", args.ConfigFile, err)
		os.Exit(1)
	}
	for {
		t, ok := f.Next()
		if !ok {
			break
		}
		fmt.Printf("[%s]\n", t.Name)
		for {
			dir, ok := t.Next()
			if !ok {
				break
			}
			fmt.Printf("\t[%s]\n", dir)
			fmt.Printf("\t\tLoading package.yml...")
			pkg, err := ypkg.Load(dir)
			if err != nil {
				fmt.Println("FAILED!")
				fmt.Printf("\t\t\t%s\n", err)
				os.Exit(1)
			}
			fmt.Println("OK.")
			fmt.Printf("\t\tFound '%s', release '%d'.\n", pkg.Name, pkg.Release)

			fmt.Printf("\t\tFetching index...")
			i, err := eopkg.FetchIndex()
			if err != nil {
				fmt.Println("FAILED!")
				fmt.Printf("\t\t\t%s\n", err)
				os.Exit(1)
			}
			fmt.Println("OK.")
			if i.HasRelease(pkg.Name, pkg.Release) {
				fmt.Printf("\t\tSkipping completed.\n")
				continue
			}
			fmt.Printf("\t\tPublishing changes...DONE\n")
		}
		fmt.Printf("[Done]\n\n")
	}
}
