package cli

import (
	"dev.getsol.us/sources/infrastructure-tooling/autopush/order"
	"fmt"
	"github.com/DataDrake/cli-ng/cmd"
	"path/filepath"
	"os"
)

var Bootstrap = cmd.Sub{
	Name: "bootstrap",
	Alias: "bt",
	Short: "Bootstrap an example autopush configuration file",
	Flags: &BootstrapFlags{},
	Args: &BootstrapArgs{},
	Run: RunBootstrap,
}

func init() {
	cmd.Register(&Bootstrap)
}

type BootstrapArgs struct{}
type BootstrapFlags struct{}

func RunBootstrap(r *cmd.Root, c *cmd.Sub) {
	defaultTier := order.Tier{
		Name: "Tier 0",
		Packages: []string{"i-am-a-directory-for-a-package"},
	}

	newFile := order.File{
		Tiers: []*order.Tier{ &defaultTier, },
	}

	workDir, getWdErr := os.Getwd()

	if getWdErr != nil { // Failed to get the working directory
		fmt.Printf("failed to get working directory: %s\n", getWdErr)
		os.Exit(1)
	}

	marshalledFile, marshalErr := newFile.Marshal() // Marshal our new file into YAML

	if marshalErr != nil { // Failed to marshal the file
		fmt.Printf("failed to marshal our example file: %s\n", marshalErr)
		os.Exit(1)
	}

	autopushPath := filepath.Join(workDir, "autopush.yml")
	autopushFile, openErr := os.OpenFile(autopushPath, os.O_RDWR|os.O_CREATE, 0755)

	if openErr != nil { // Failed to open the file
		fmt.Printf("failed to open the file: %s", openErr)
		os.Exit(1)
	}

	defer autopushFile.Close() // Defer closing
	if _, writeErr := autopushFile.Write(marshalledFile); writeErr == nil { // Attempt to write our marshalled file contents
		fmt.Printf("wrote new autopush file to: %s\n", autopushPath)
	} else { // Failed to write the file
		fmt.Printf("failed to write the file: %s\n", writeErr)
		os.Exit(1)
	}
}