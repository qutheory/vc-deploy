package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strings"
	"syscall"
	"github.com/qutheory/vc-deploy/api"
)

var (
	loginCmd = &cobra.Command{
		Use:   "login",
		Short: "Login to Vapor Cloud 2",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			reader := bufio.NewReader(os.Stdin)

			fmt.Print("Enter Email: ")
			enteredEmail, _ := reader.ReadString('\n')

			fmt.Print("Enter Password: ")
			bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))

			enteredPassword := string(bytePassword)

			email := strings.TrimSpace(enteredEmail)
			password := strings.TrimSpace(enteredPassword)

			_, err := api.Login(email, password)
			fmt.Println("")
			if err != nil {
				fmt.Println("Authentication failed")
				os.Exit(1)
			}
			fmt.Println("You are now signed in.")
		},
	}
)

func init() {
	rootCmd.AddCommand(loginCmd)
}
