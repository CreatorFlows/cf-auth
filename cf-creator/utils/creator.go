package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/hanshal101/cf-creator/database/models"
	"github.com/hanshal101/cf-creator/logger"
	"github.com/hanshal101/cf-creator/utils"
	"go.uber.org/zap"
)

func GetCreator(c *gin.Context) (models.Owner, error) {
	token, err := c.Cookie("tokenCF")
	if err != nil {
		logger.Logger.Warn("error in finding token", zap.Error(err))
		return models.Owner{}, err
	}

	_, err = utils.ParseToken(token)
	if err != nil {
		logger.Logger.Warn("error in parsing token", zap.Error(err))
		return models.Owner{}, err
	}
	return models.Owner{}, nil
}
