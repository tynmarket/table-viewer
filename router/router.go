package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tynmarket/table-viewer/controller"
)

// Route is
func Route(r *gin.Engine) *gin.Engine {
	r.POST("/select", controller.Select)

	return r
}