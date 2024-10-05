package cmd

import (
	"cognito-token-generator-cli/internal"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// Root command definition
var rootCmd = &cobra.Command{
	Use:   "jwtcli",
	Short: "JWT CLI Tool to generate AWS Cognito JWT tokens",
	Run: func(cmd *cobra.Command, args []string) {
		var creds internal.Credentials

		// Prompt the user for an encryption passphrase
		passphrase := internal.ReadInput("Enter the encryption passphrase: ")

		// Derive the encryption key from the passphrase
		encryptionKey := internal.DeriveEncryptionKey(passphrase)

		// Check if credentials file exists
		if internal.FileExists(internal.GetCredentialsFilePath()) {
			fmt.Println("Found saved AWS credentials, loading them...")

			// Load saved credentials
			savedCreds, err := internal.LoadCredentials(internal.GetCredentialsFilePath())
			if err != nil {
				log.Fatalf("Error loading saved credentials: %v", err)
			}
			creds = *savedCreds

			// Validate if ClientID, UserPoolID, and Region are not empty
			if creds.ClientID == "" || creds.UserPoolID == "" || creds.Region == "" {
				fmt.Println("AWS credentials are incomplete, please enter them.")
				creds.ClientID = internal.ReadInput("Enter Cognito Client ID: ")
				creds.UserPoolID = internal.ReadInput("Enter Cognito User Pool ID: ")
				creds.Region = internal.ReadInput("Enter AWS Region (e.g., us-west-0): ")

				// Save updated credentials
				err := internal.SaveCredentials(&creds, internal.GetCredentialsFilePath())
				if err != nil {
					log.Fatalf("Failed to save credentials: %v", err)
				}
			} else {
				fmt.Printf("\nThe following saved AWS credentials are being used:\n")
				fmt.Printf("Client ID: %s\n", creds.ClientID)
				fmt.Printf("User Pool ID: %s\n", creds.UserPoolID)
				fmt.Printf("Region: %s\n", creds.Region)
			}
		} else {
			// Prompt user for credentials if they are not saved
			creds.ClientID = internal.ReadInput("Enter Cognito Client ID: ")
			creds.UserPoolID = internal.ReadInput("Enter Cognito User Pool ID: ")
			creds.Region = internal.ReadInput("Enter AWS Region (e.g., us-west-0): ")

			// Ask user if they want to save credentials
			save := internal.ReadInput("Do you want to save these credentials for future use? (yes/no): ")
			if save == "yes" {
				internal.SaveCredentials(&creds, internal.GetCredentialsFilePath())
			}
		}

		// Choose user or add new credentials, passing the encryption key
		selectedUser, err := internal.SelectUserOrNew(&creds, encryptionKey)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}

		// Decrypt password using the derived encryption key
		password, err := internal.Decrypt(selectedUser.EncryptedPassword, encryptionKey)
		if err != nil {
			log.Fatalf("Error decrypting password: %v", err)
		}

		// Save credentials with the newly added user
		err = internal.SaveCredentials(&creds, internal.GetCredentialsFilePath())
		if err != nil {
			log.Fatalf("Failed to save credentials: %v", err)
		}

		// Load AWS SDK Config with saved or entered region
		cfg, err := internal.LoadAWSConfig(creds.Region)
		if err != nil {
			log.Fatalf("unable to load AWS SDK config: %v", err)
		}

		// Call AWS Cognito to authenticate user and get JWT
		idToken, err := internal.AuthenticateWithCognito(cfg, creds.ClientID, selectedUser.Username, password)
		if err != nil {
			log.Fatalf("Failed to authenticate user: %v", err)
		}

		// Output JWT token and copy to clipboard
		fmt.Printf("\nJWT Token: %s\n", idToken)
		err = internal.CopyToClipboard(idToken)
		if err != nil {
			log.Fatalf("Failed to copy JWT to clipboard: %v", err)
		}
		fmt.Println("\nJWT token has been copied to clipboard.")
	},
}

// Execute runs the root command
func Execute() {
	rootCmd.AddCommand(clearAWSCredsCmd)
	rootCmd.AddCommand(clearUserCredsCmd)
	rootCmd.AddCommand(versionCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
