package cmd

import (
	"fmt"
	"os"

	"github.com/qutheory/vc-deploy/api"
	"github.com/spf13/cobra"
)

var (
	deployCmd = &cobra.Command{
		Use:   "deploy",
		Short: "Deploy your app",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Deploying: " + Slug + " Environment: " + Env)

			_, env, err := api.GetAppEnv(Slug, Env, Token)
			if err != nil {
				fmt.Println("App/Env not found")
				os.Exit(1)
			}

			activity, err := api.Deploy(env.Id, Branch, Token)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			api.Listen(activity.Activity.Id)
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
	deployCmd.Flags().StringVarP(&Token, "token", "t", "", "Developer token")
	deployCmd.Flags().StringVarP(&Branch, "branch", "b", "", "Branch")
	deployCmd.Flags().BoolVarP(&NoFollow, "no-follow", "", false, "Don't follow the log output")

	deployCmd.MarkFlagRequired("app")
	deployCmd.MarkFlagRequired("env")
	deployCmd.MarkFlagRequired("token")
}
