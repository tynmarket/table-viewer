package model

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // blank import
)

var user = os.Getenv("DB_USER")
var password = os.Getenv("DB_PASSWORD")
var host = os.Getenv("DB_HOST")
var database = os.Getenv("DB_NAME")
var url string

// Model is
type Model struct {
	ID        int64     `gorm:"primary_key" json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// Db is
func Db() (*gorm.DB, error) {
	return gorm.Open("mysql", url)
}

// Records is
func Records(rows *sql.Rows) ([][]string, error) {
	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	rawResult := make([][]byte, len(cols))
	dest := make([]interface{}, len(cols))
	records := make([][]string, 0, 1000)

	for i := range rawResult {
		dest[i] = &rawResult[i]
	}

	for rows.Next() {
		result := make([]string, len(cols))

		err = rows.Scan(dest...)
		if err != nil {
			fmt.Println("Failed to scan row", err)
			return nil, err
		}

		for i, raw := range rawResult {
			if raw == nil {
				result[i] = "NULL"
			} else {
				result[i] = string(raw)
			}
		}

		records = append(records, result)

		fmt.Printf("%#v\n", result)
	}

	return records, nil
}

func init() {
	if host == "" {
		host = "127.0.0.1"
	}

	url = fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, database)
}
