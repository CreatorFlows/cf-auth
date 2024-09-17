package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Error in initiating env: %v\n", err)
		os.Exit(1)
	}
}

func main() {
	r := gin.New()

	r.Run(os.Getenv("APP_PORT"))
}
