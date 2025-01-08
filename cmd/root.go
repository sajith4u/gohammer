package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gohammer",
	Short: "A load testing tool for APIs",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use 'gohammer run' to start a test.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
