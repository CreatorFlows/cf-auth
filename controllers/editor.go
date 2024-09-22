package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/creatorflows/cf-auth/config"
	"github.com/creatorflows/cf-auth/database/models"
	"github.com/creatorflows/cf-auth/database/postres"
	"github.com/creatorflows/cf-auth/logger"
	"github.com/creatorflows/cf-auth/utils"
)

func Editor(c *gin.Context) {
	cookie, err := c.Cookie("tokenCF")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	claims, err := utils.ParseToken(cookie)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if claims.Role != "EDITOR" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "editor page", "role": claims.Role})
}

func EditorLogin(c *gin.Context) {
	var editor models.Editor
	if err := c.ShouldBindJSON(&editor); err != nil {
		logger.Logger.Warn("error in binding", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "error in binding request"})
		return
	}

	var existingUser models.Editor
	if err := postres.DB.Where("email = ?", editor.Email).First(&existingUser).Error; err != nil {
		logger.Logger.Warn("error in finding user", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error in finding user"})
		return
	}

	if existingUser.ID == 0 {
		logger.Logger.Info("user does not exist")
		c.JSON(http.StatusBadRequest, gin.H{"error": "user does not exist"})
		return
	}

	// errHash := utils.CompareHashPassword(editor.Password, existingUser.Password)
	// if !errHash {
	// 	logger.Logger.Info("invalid password")
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "invalid password"})
	// 	return
	// }
	if editor.Password != existingUser.Password {
		logger.Logger.Info("invalid password")
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid password"})
		return
	}
	tokenString, err := utils.CreateClaims("EDITOR", existingUser.Email, config.EXP_TIME.Unix())
	if err != nil {
		logger.Logger.Warn("error in creating jwt token", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error in creating token"})
		return
	}

	c.SetCookie("tokenCF", tokenString, int(config.EXP_TIME.Unix()), "/", "", false, false)
	c.JSON(http.StatusOK, gin.H{"success": tokenString})
}

func EditorLogout(c *gin.Context) {
	c.SetCookie("tokenCF", "", -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"success": "user logged out"})
}
