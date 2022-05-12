package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func configEnv() error {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			fmt.Println("Could not load .env file")
			return err
		}
	}
	return nil
}
