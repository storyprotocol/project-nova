package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// dbCmd represents the db command
var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "db task runs db operation",
	Long:  `db task runs db operation. To be implemented`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("db called")
	},
}

func init() {
	rootCmd.AddCommand(dbCmd)
}
