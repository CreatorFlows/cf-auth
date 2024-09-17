package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hanshal101/cf-auth/config"
	"github.com/hanshal101/cf-auth/database/migrate"
	"github.com/hanshal101/cf-auth/database/postres"
	"github.com/hanshal101/cf-auth/logger"
	"github.com/hanshal101/cf-auth/routes"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Error in intializing env: %v\n", err)
		os.Exit(1)
	}
	logger.InitLogger()
	postres.InitDB()
	migrate.AutoMigrate()
	config.JWT_KEY = []byte(os.Getenv("JWT_KEY"))
}

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	routes.AuthRoutes(r)
	r.Run(os.Getenv("APP_PORT"))
}
