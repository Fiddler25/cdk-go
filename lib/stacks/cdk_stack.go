package stacks

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvNames struct {
	KeyName string
}

func CdkEnvNames() EnvNames {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	return EnvNames{
		KeyName: os.Getenv("KEY_NAME"),
	}
}
