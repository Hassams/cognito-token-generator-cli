package internal

import (
	"encoding/json"
	"os"
)

// Credentials holds AWS and user credentials
type Credentials struct {
	ClientID   string `json:"clientID"`
	UserPoolID string `json:"userPoolID"`
	Region     string `json:"region"`
	Users      []User `json:"users"`
}

// User holds username and encrypted password
type User struct {
	Username          string `json:"username"`
	EncryptedPassword string `json:"encryptedPassword"`
}

// SaveCredentials saves the credentials to a file
func SaveCredentials(credentials *Credentials, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(credentials)
}

// LoadCredentials loads the credentials from a file
func LoadCredentials(path string) (*Credentials, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var credentials Credentials
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&credentials)
	if err != nil {
		return nil, err
	}
	return &credentials, nil
}
