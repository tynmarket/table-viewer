package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/tynmarket/table-viewer/model"
)

func handle(c *gin.Context, handlerFun func(db *gorm.DB)) {
	db, err := model.Db()

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"message": "error",
		})
		return
	}

	defer db.Close()
	db.LogMode(true)

	handlerFun(db)
}
