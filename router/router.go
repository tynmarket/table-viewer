package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tynmarket/table-viewer/controller"
)

// Route is
func Route(r *gin.RouterGroup) {
	r.POST("/select", controller.Select)
}
