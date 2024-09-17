package postres

import (
	"os"

	"github.com/hanshal101/cf-auth/logger"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := os.Getenv("POSTGRES_DB_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Logger.Error("Error in loading the DB: %v\n", zap.Error(err))
		return
	}
	DB = db
	logger.Logger.Info("Postgres Initialized !!!")
}
