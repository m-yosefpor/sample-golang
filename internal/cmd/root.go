package cmd

import (
	"github.com/spf13/cobra"
)

const (
	use   = "httpmon"
	short = "httpmon app"
)

func NewRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   use,
		Short: short,
	}
	rootCmd.AddCommand(newVersionCmd())
	rootCmd.AddCommand(newStartCmd())
	return rootCmd
}
