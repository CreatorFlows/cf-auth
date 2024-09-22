package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/creatorflows/cf-auth/controllers"
	"github.com/creatorflows/cf-auth/middlewares"
)

func AuthRoutes(r *gin.Engine) {
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.Use(middlewares.IsAuthorized())
	r.GET("/home", controllers.Home)
	r.GET("/premium", controllers.Premium)
	r.POST("/conv/premium", controllers.ConverttoPremium)
	r.GET("/logout", controllers.Logout)
}

func EditorAuthRoutes(r *gin.Engine) {
	r.GET("/editor", controllers.Editor)
	r.POST("/editor/login", controllers.EditorLogin)
	r.GET("/editor/logout", controllers.EditorLogout)
}
