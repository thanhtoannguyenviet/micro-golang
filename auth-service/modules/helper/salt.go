package helper

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

const saltSize = 16

func GenerateRandomSalt() string {
	var salt = make([]byte, saltSize)
	_, err := rand.Read(salt[:])
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(salt)
}
func HashPassword(password, salt string) string {
	var passwordBytes = []byte(password)
	var sha512Hasher = sha256.New()
	passwordBytes = append(passwordBytes, salt...)
	sha512Hasher.Write(passwordBytes)
	var hashedPasswordBytes = sha512Hasher.Sum(nil)
	var hashedPasswordHex = hex.EncodeToString(hashedPasswordBytes)
	return hashedPasswordHex
}
func ConstainsPassword(hashPassword, currentPassword, salt string) bool {
	currPasswordHash := HashPassword(currentPassword, salt)
	return hashPassword == currPasswordHash
}
