package cmd

import (
	"os"
	"os/exec"
	"syscall"

	"github.com/project-nova/backend/bastion/config"
	"github.com/project-nova/backend/pkg/logger"
	"github.com/spf13/cobra"
)

var dbShellCmd = &cobra.Command{
	Use:   "db-shell",
	Short: "db shell task opens the command interface to AWS Postgres",
	Long:  `db shell task opens the command interface to AWS Postgres`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.GetConfig()
		if err != nil {
			logger.Fatalf("Failed to get configs: %v", err)
		}

		binary, err := exec.LookPath("psql")
		if err != nil {
			logger.Fatalf("Failed to find psql path: %v", err)
		}

		cmdArgs := []string{"psql", cfg.DatabaseURI}
		env := os.Environ()

		execErr := syscall.Exec(binary, cmdArgs, env)
		if execErr != nil {
			logger.Fatalf("Failed to execute the command: %v", execErr)
		}
	},
}

var dbUpCmd = &cobra.Command{
	Use:   "db-up",
	Short: "db up task runs db migration to upgrade db to the latest version",
	Long:  `db up task runs db migration to upgrade db to the latest version`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.GetConfig()
		if err != nil {
			logger.Fatalf("Failed to get configs: %v", err)
		}

		binary, err := exec.LookPath("migrate")
		if err != nil {
			logger.Fatalf("Failed to find psql path: %v", err)
		}

		cmdArgs := []string{"migrate", "-database", cfg.DatabaseURI, "-path", "/build/api/migrations", "-verbose", "up"}
		env := os.Environ()

		execErr := syscall.Exec(binary, cmdArgs, env)
		if execErr != nil {
			logger.Fatalf("Failed to execute the command: %v", execErr)
		}
	},
}

func init() {
	rootCmd.AddCommand(dbShellCmd)
	rootCmd.AddCommand(dbUpCmd)
}
