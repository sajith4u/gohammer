package cmd

import (
	"gohammer/internal/loader"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a load test",
	Run: func(cmd *cobra.Command, args []string) {
		loader.RunLoadTest("https://example.com", 10, 100)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
