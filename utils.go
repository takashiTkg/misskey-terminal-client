package main

import (
	"fmt"

	"github.com/joho/godotenv"
)

// load .env file
func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf(".env load failed. %v", err)
	}
}
