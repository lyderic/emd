/*
Copyright © 2024 Lydéric Landry
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "emd",
	Short: "A basic markdown viewer",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			return render(args[0])
		}
		return fmt.Errorf("missing argument!")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
