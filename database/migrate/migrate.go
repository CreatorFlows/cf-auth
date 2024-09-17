package migrate

import (
	"github.com/hanshal101/cf-auth/database/models"
	"github.com/hanshal101/cf-auth/database/postres"
	"github.com/hanshal101/cf-auth/logger"
)

func AutoMigrate() {
	tx := postres.DB.Begin()
	tx.AutoMigrate(&models.Owner{})
	tx.AutoMigrate(&models.Editor{})
	tx.Commit()
	logger.Logger.Info("Models Migrated !!!")
}
