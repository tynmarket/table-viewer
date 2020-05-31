package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type selectParam struct {
	Table        string       `json:"table" binding:"required"`
	WhereQueries []whereQuery `json:"where"`
	Order        string       `json:"order"`
}

type whereQuery struct {
	Column   string `json:"column" binding:"required"`
	Operator string `json:"operator" binding:"required"`
	Value    string `json:"value"`
}

// Select is
func Select(c *gin.Context) {
	handle(c, func(db *gorm.DB) {
		var json selectParam
		if err := c.ShouldBindJSON(&json); err != nil {
			fmt.Printf("error: %s", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Printf("param: %v", json)
		c.JSON(http.StatusOK, gin.H{"param": json})
	})
}
