package migrate

import (
	"github.com/hanshal101/cf-creator/database/models"
	"github.com/hanshal101/cf-creator/database/postgres"
	"github.com/hanshal101/cf-creator/logger"
)

func AutoMigrate() {
	tx := postgres.DB.Begin()
	tx.AutoMigrate(&models.Owner{})
	tx.AutoMigrate(&models.Editor{})
	tx.Commit()
	logger.Logger.Info("Models Migrated !!!")

}
