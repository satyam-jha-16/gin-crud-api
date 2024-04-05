package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error while loading env file")
	} else {
		log.Print("env loaded successfully")
	}
}
