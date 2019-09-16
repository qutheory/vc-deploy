package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	testCmd = &cobra.Command{
		Use:   "test",
		Short: "Simple test command",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("The CLI works! :)")
		},
	}
)

func init() {
	rootCmd.AddCommand(testCmd)
}
