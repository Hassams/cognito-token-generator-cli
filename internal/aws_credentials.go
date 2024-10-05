package internal

import (
	"os"
	"path/filepath"
)

// ClearAWSCredentials clears the saved AWS credentials
func ClearAWSCredentials() error {
	creds, err := LoadCredentials(GetCredentialsFilePath())
	if err != nil {
		return err
	}
	creds.ClientID = ""
	creds.UserPoolID = ""
	creds.Region = ""

	return SaveCredentials(creds, GetCredentialsFilePath())
}

// GetCredentialsFilePath returns the file path for the credentials
func GetCredentialsFilePath() string {
	homeDir, _ := os.UserHomeDir()
	return filepath.Join(homeDir, ".aws_cognito_credentials.json")
}
