package migrate

import (
	"github.com/creatorflows/cf-auth/database/models"
	"github.com/creatorflows/cf-auth/database/postres"
	"github.com/creatorflows/cf-auth/logger"
)

func AutoMigrate() {
	tx := postres.DB.Begin()
	tx.AutoMigrate(&models.Owner{})
	tx.AutoMigrate(&models.Editor{})
	tx.Commit()
	logger.Logger.Info("Models Migrated !!!")
}
