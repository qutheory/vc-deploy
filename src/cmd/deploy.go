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
			fmt.Println("Deploying " + Slug)

			if !NoFollow {
				fmt.Println("Follow!")
			}

			// api.GetUser()
		},
	}
)

var Slug string
var Env string
var Branch string
var Token string
var NoFollow bool

func init() {
	rootCmd.AddCommand(deployCmd)

	deployCmd.Flags().StringVarP(&Slug, "app", "a", "", "Application slug")
	deployCmd.Flags().StringVarP(&Env, "env", "e", "", "Environment")
	deployCmd.Flags().StringVarP(&Slug, "token", "t", "", "Developer token")
	deployCmd.Flags().StringVarP(&Slug, "branch", "b", "", "Branch")
	deployCmd.Flags().BoolVarP(&NoFollow, "no-follow", "", false, "Don't follow the log output")
}
