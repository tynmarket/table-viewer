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

type responseData struct {
	Header  []string   `json:"header"`
	Records [][]string `json:"records"`
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

		columnTypes, err := rows.ColumnTypes()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"rows to records": err.Error()})
		}

		fmt.Printf("\ncolumnType: %#v\n", columnTypes[0])

		records, err := model.Records(rows)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"rows to records": err.Error()})
		}

		fmt.Printf("\n%d rows returned\n\n", len(records))

		var header []string

		for _, columnType := range columnTypes {
			header = append(header, columnType.Name())
		}

		data := &responseData{
			Header:  header,
			Records: records,
		}

		c.JSON(http.StatusOK, gin.H{"data": data})
	})
}
