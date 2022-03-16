package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const VERSION = "0.1.0"

func newVersionCmd() *cobra.Command {

	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "print version",
		Run: func(c *cobra.Command, args []string) {
			version()
		},
	}
	return versionCmd
}

func version() {
	fmt.Println(VERSION)
}
