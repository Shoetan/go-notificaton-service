package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)


func GetEnvVariables(keys ...string) map[string]string {
	err := godotenv.Load()
	if err != nil {
			log.Printf("Error loading .env: %s", err.Error())
	}

	envVars := make(map[string]string)
	for _, key := range keys {
			envVars[key] = os.Getenv(key)
	}

	return envVars
}

func HashPwd( password string) (string, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hash), err
	
}
