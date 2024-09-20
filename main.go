package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/creatorflows/cf-auth/config"
	"github.com/creatorflows/cf-auth/database/migrate"
	"github.com/creatorflows/cf-auth/database/postres"
	"github.com/creatorflows/cf-auth/logger"
	"github.com/creatorflows/cf-auth/routes"
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
