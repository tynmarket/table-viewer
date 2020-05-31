package model

import (
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"
)

var host = os.Getenv("HOST")
var database = os.Getenv("DATABASE")
var user = os.Getenv("USER")
var password = os.Getenv("PASSWORD")
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

func init() {
	if host == "" {
		host = "127.0.0.1"
	}

	url = fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, database)
}
