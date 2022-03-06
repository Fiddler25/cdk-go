package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	KeyName            string
	MasterUserPassword string
}

func EnvNames() Env {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	return Env{
		KeyName:            os.Getenv("KEY_NAME"),
		MasterUserPassword: os.Getenv("MASTER_USER_PASSWORD"),
	}
}
