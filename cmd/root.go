package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "0.1.0"

var rootCmd = &cobra.Command{
	Use:   "aplookup",
	Short: "A simple CLI to look up with AcitivityPub",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("aplookup v%s\n", version)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
