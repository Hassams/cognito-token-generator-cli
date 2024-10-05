package cmd

import (
	"cognito-token-generator-cli/internal"
	"fmt"
	"github.com/spf13/cobra"
)

var clearAWSCredsCmd = &cobra.Command{
	Use:   "clear-aws-credentials",
	Short: "Clear saved AWS Cognito credentials",
	Run: func(cmd *cobra.Command, args []string) {
		err := internal.ClearAWSCredentials()
		if err != nil {
			fmt.Println("Error clearing AWS credentials:", err)
		} else {
			fmt.Println("AWS credentials cleared.")
		}
	},
}

var clearUserCredsCmd = &cobra.Command{
	Use:   "clear-user-credentials",
	Short: "Clear saved user credentials",
	Run: func(cmd *cobra.Command, args []string) {
		err := internal.ClearUserCredentials()
		if err != nil {
			fmt.Println("Error clearing user credentials:", err)
		} else {
			fmt.Println("User credentials cleared.")
		}
	},
}
