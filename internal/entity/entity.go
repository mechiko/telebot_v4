package entity

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var Mode = ""

func init() {
	// loads values from .env into the system
	if os.Getenv("MODE") == "" {
		if err := godotenv.Load(); err != nil {
			fmt.Println("entity .env not found")
		}
	}
	if value, exists := os.LookupEnv("MODE"); exists {
		Mode = value
	}
}
