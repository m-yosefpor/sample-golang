package cmd

import (
	"github.com/spf13/cobra"
)

func newMigrateCmd() *cobra.Command {

	migrateCmd := &cobra.Command{
		Use:   "migrate",
		Short: "migrate the db",
		Run: func(cmd *cobra.Command, args []string) {
			migrate()
		},
	}
	return migrateCmd
}

func migrate() {}
