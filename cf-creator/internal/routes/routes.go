package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hanshal101/cf-creator/internal/editors"
)

func EditorRoutes(r *gin.Engine) {
	r.GET("", editors.GetEditors)
}
