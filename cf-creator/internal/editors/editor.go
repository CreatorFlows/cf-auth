package editors

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hanshal101/cf-creator/database/models"
	"github.com/hanshal101/cf-creator/database/postgres"
	"github.com/hanshal101/cf-creator/logger"
	"go.uber.org/zap"
)

func GetEditors(c *gin.Context) {
	var editors []models.Editor
	if err := postgres.DB.Find(&editors).Error; err != nil {
		logger.Logger.Warn("error in fetching editors", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error in fetching editors"})
		return
	}
	c.JSON(http.StatusOK, editors)
}

func CreateEditors(c *gin.Context) {
	var editor models.Editor
	if err := c.ShouldBindJSON(&editor); err != nil {
		logger.Logger.Warn("error in binding request", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "error in binding request"})
		return
	}

	if err := postgres.DB.Create(&editor).Error; err != nil {
		logger.Logger.Warn("error in creating editor", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error in creating editor"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "editor created"})
}

func DeleteEditor(c *gin.Context) {
	editorID := c.Param("editorID")
	if err := postgres.DB.Where("id = ?", editorID).Delete(&models.Editor{}).Error; err != nil {
		logger.Logger.Warn("error in deleting editor", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error in deleting editor"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "editor deleted"})
}

func EditEditor(c *gin.Context) {
	var editor models.Editor
	if err := c.ShouldBindJSON(&editor); err != nil {
		logger.Logger.Warn("error in binding request", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "error in binding request"})
		return
	}

	var existingEditor models.Editor
	if err := postgres.DB.Find(&existingEditor).Error; err != nil {
		logger.Logger.Warn("error in fetching editors", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error in fetching editors"})
		return
	}

	existingEditor.Name = editor.Name
	existingEditor.Phone = editor.Phone
	if err := postgres.DB.Find(&editor).Error; err != nil {
		logger.Logger.Warn("error in fetching editors", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error in fetching editors"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "editor edited"})
}
