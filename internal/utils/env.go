package utils

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var once sync.Once

func GetEnv(envPath string) string {
	once.Do(func() {
		if err := godotenv.Load(); err != nil {
			log.Fatal(err)
		}
	})

	secret, ok := os.LookupEnv(envPath)
	if !ok {
		log.Println("Failed getting env value")
		return secret
	}

	if secret == "" {
		log.Println("Secret value is empty")
		return secret
	}

	return secret
}
