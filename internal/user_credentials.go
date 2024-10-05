package internal

import "fmt"

func ClearUserCredentials() error {
	creds, err := LoadCredentials(GetCredentialsFilePath())
	if err != nil {
		return err
	}
	creds.Users = []User{}
	return SaveCredentials(creds, GetCredentialsFilePath())
}

// SelectUserOrNew allows the user to select an existing user or enter new credentials
func SelectUserOrNew(creds *Credentials, encryptionKey []byte) (*User, error) {
	if len(creds.Users) > 0 {
		fmt.Println("Do you want to use an existing user or enter new credentials?")
		fmt.Println("1. Choose from saved users")
		fmt.Println("2. Enter new credentials")
		choice := ReadInput("Enter choice (1/2): ")

		if choice == "1" {
			fmt.Println("Available users:")
			for i, user := range creds.Users {
				fmt.Printf("%d. %s\n", i+1, user.Username)
			}
			selectedUser := ReadInput("Select a user by number: ")
			userIndex := int(selectedUser[0] - '1')
			if userIndex < 0 || userIndex >= len(creds.Users) {
				return nil, fmt.Errorf("invalid selection")
			}
			return &creds.Users[userIndex], nil
		}
	}

	// If no users or user chooses to enter new credentials
	username := ReadInput("Enter username: ")
	password := ReadInput("Enter password: ")

	// Encrypt the password using the dynamically derived encryption key
	encryptedPassword, err := Encrypt(password, encryptionKey)
	if err != nil {
		return nil, err
	}

	newUser := User{
		Username:          username,
		EncryptedPassword: encryptedPassword,
	}

	// Add to saved users
	creds.Users = append(creds.Users, newUser)
	return &newUser, nil
}
