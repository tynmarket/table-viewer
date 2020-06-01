package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/tynmarket/table-viewer/model"
)

type selectParam struct {
	Table        string             `json:"table" binding:"required"`
	WhereQueries []model.WhereQuery `json:"where"`
	Order        string             `json:"order"`
}

// Select is
func Select(c *gin.Context) {
	handle(c, func(db *gorm.DB) {
		var param selectParam
		if err := c.ShouldBindJSON(&param); err != nil {
			fmt.Printf("error: %s", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Printf("\nparam: %#v\n", param)

		rows, err := model.
			BuildQuery(db, param.Table, param.WhereQueries, param.Order).
			Select("*").
			Rows()

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"select error": err.Error()})
		}

		records, err := model.Records(rows)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"rows to records": err.Error()})
		}

		fmt.Printf("\n%d rows returned\n\n", len(records)-1)

		c.JSON(http.StatusOK, gin.H{"data": records})
	})
}
