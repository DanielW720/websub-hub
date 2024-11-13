package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"hub/src/models"
	"math/rand"
)

// Generate a random string of a given length
func GenerateRandomString(length int) string {
	letters := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	bytes := make([]byte, length)

	for i := range bytes {
		bytes[i] = letters[rand.Intn(len(letters))]
	}

	return string(bytes)
}

// Generate a random name and surname
func GenerateUser() models.User {
	firstNames := []string{"John", "Jane", "Alex", "Emily", "Chris", "Katie"}
	lastNames := []string{"Smith", "Johnson", "Williams", "Brown", "Jones", "Garcia"}

	firstName := firstNames[rand.Intn(len(firstNames))]
	lastName := lastNames[rand.Intn(len(lastNames))]
	age := rand.Intn(100) // Random age between 0 and 99

	user := models.User{
		Firstname: firstName,
		Lastname:  lastName,
		Age:       age,
	}

	return user
}

// Generates an HMAC hash using the given key and message
func HashWithHMAC(message, key []byte) string {
	h := hmac.New(sha256.New, key)
	h.Write(message)
	return hex.EncodeToString(h.Sum(nil))
}
