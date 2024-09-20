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

func Login(c *gin.Context) {

	var user models.Owner
	if err := c.ShouldBindJSON(&user); err != nil {
		logger.Logger.Warn("error in binding", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "error in binding request"})
		return
	}

	var existingUser models.Owner
	if err := postres.DB.Where("email = ?", user.Email).First(&existingUser).Error; err != nil {
		logger.Logger.Warn("error in finding user", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error in finding user"})
		return
	}

	if existingUser.ID == 0 {
		logger.Logger.Info("user does not exist")
		c.JSON(http.StatusBadRequest, gin.H{"error": "user does not exist"})
		return
	}

	errHash := utils.CompareHashPassword(user.Password, existingUser.Password)
	if !errHash {
		logger.Logger.Info("invalid password")
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid password"})
		return
	}

	tokenString, err := utils.CreateClaims(existingUser.Role, existingUser.Email, config.EXP_TIME.Unix())
	if err != nil {
		logger.Logger.Warn("error in creating jwt token", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error in creating token"})
		return
	}

	c.SetCookie("tokenCF", tokenString, int(config.EXP_TIME.Unix()), "/", "", false, false)
	c.JSON(http.StatusOK, gin.H{"success": tokenString})
}

func Signup(c *gin.Context) {
	var user models.Owner
	if err := c.ShouldBindJSON(&user); err != nil {
		logger.Logger.Warn("error in binding", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingUser models.Owner
	if err := postres.DB.Where("email = ?", user.Email).Find(&existingUser).Error; err != nil {
		logger.Logger.Warn("error in fetching user", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if existingUser.ID != 0 {
		logger.Logger.Info("user already exists")
		c.JSON(http.StatusBadRequest, gin.H{"error": "user already exists"})
		return
	}

	var errHash error
	user.Password, errHash = utils.GenerateHashPassword(user.Password)
	if errHash != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "could not generate password hash"})
		return
	}

	user.Role = "GUEST"

	if err := postres.DB.Create(&user).Error; err != nil {
		logger.Logger.Warn("error in creating user", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "user created"})

}

func Logout(c *gin.Context) {
	c.SetCookie("tokenCF", "", -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"success": "user logged out"})
}

func Home(c *gin.Context) {
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

	if claims.Role != "GUEST" && claims.Role != "CREATOR" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "home page", "role": claims.Role})
}

func Premium(c *gin.Context) {
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
	if claims.Role != "CREATOR" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": "premium page", "role": claims.Role})
}

func ConverttoPremium(c *gin.Context) {
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

	if claims.Role != "GUEST" {
		if claims.Role == "CREATOR" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "already a CREATOR"})
			return
		}
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var creator models.Owner
	if err := postres.DB.Where("email = ?", claims.Subject).Find(&creator).Error; err != nil {
		logger.Logger.Warn("error in fetching user", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := postres.DB.Model(&creator).Update("role", "CREATOR").Error; err != nil {
		logger.Logger.Warn("error in updating user", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokenString, err := utils.CreateClaims(creator.Role, creator.Email, config.EXP_TIME.Unix())
	if err != nil {
		logger.Logger.Warn("error in creating jwt token", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error in creating token"})
		return
	}

	c.SetCookie("tokenCF", tokenString, int(config.EXP_TIME.Unix()), "/", "", false, false)
}
