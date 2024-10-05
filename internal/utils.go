package internal

import (
	"bufio"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"os"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

// FileExists checks if a file exists at the specified path
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// ReadInput prompts the user for input and reads it from stdin
func ReadInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// LoadAWSConfig loads the AWS SDK configuration using the specified region
func LoadAWSConfig(region string) (aws.Config, error) {
	return config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
}

// AuthenticateWithCognito authenticates a user with AWS Cognito and returns the Access Token
func AuthenticateWithCognito(cfg aws.Config, clientID, username, password string) (string, error) {
	cognitoClient := cognitoidentityprovider.NewFromConfig(cfg)

	input := &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: "USER_PASSWORD_AUTH", // Authentication flow for username & password
		AuthParameters: map[string]string{
			"USERNAME": username,
			"PASSWORD": password,
		},
		ClientId: aws.String(clientID),
	}

	result, err := cognitoClient.InitiateAuth(context.TODO(), input)
	if err != nil {
		return "", err
	}

	// Return the Access Token (not the ID Token)
	accessToken := result.AuthenticationResult.AccessToken
	return *accessToken, nil
}

// CopyToClipboard copies the provided JWT token to the system clipboard
func CopyToClipboard(jwtToken string) error {
	return clipboard.WriteAll(jwtToken)
}
