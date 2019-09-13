package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	deployCmd = &cobra.Command{
		Use:   "deploy",
		Short: "Deploy your app",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("DEPLOY!!")
		},
	}
)

func init() {
	rootCmd.AddCommand(deployCmd)
}
