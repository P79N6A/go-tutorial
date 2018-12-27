package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	name string
	age  int
)

// root command
var rootCmd = &cobra.Command{
	Use:   "demo",
	Short: "A test demo",
	Long:  "A demo application to demonstrate usage of cobra",
	Run: func(cmd *cobra.Command, args []string) {
		if len(name) == 0 {
			cmd.Help()
			return
		}
		fmt.Println(name, age)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// command arg --flag
func init() {
	rootCmd.Flags().StringVarP(&name, "name", "n", "", "person's name")
	rootCmd.Flags().IntVarP(&age, "age", "a", 0, "person's age")
}
