package internal

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"golang.org/x/crypto/pbkdf2"
)

// DeriveEncryptionKey derives a 32-byte key from the given passphrase using PBKDF2
func DeriveEncryptionKey(passphrase string) []byte {
	salt := []byte("some-random-salt") // Use a random salt or a predefined salt
	return pbkdf2.Key([]byte(passphrase), salt, 4096, 32, sha256.New)
}

// Encrypt encrypts the plain text using the derived key
func Encrypt(plainText string, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	plainTextBytes := []byte(plainText)
	cfb := cipher.NewCFBEncrypter(block, key[:aes.BlockSize])
	cipherText := make([]byte, len(plainTextBytes))
	cfb.XORKeyStream(cipherText, plainTextBytes)
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

// Decrypt decrypts the cipher text using the derived key
func Decrypt(cipherText string, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	cipherTextBytes, _ := base64.StdEncoding.DecodeString(cipherText)
	cfb := cipher.NewCFBDecrypter(block, key[:aes.BlockSize])
	plainText := make([]byte, len(cipherTextBytes))
	cfb.XORKeyStream(plainText, cipherTextBytes)
	return string(plainText), nil
}
