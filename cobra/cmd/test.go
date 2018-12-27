package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test ",
	Short: "short for test command",
	Long:  "long for test command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}
