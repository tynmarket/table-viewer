package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tynmarket/table-viewer/router"
)

var authUser = os.Getenv("AUTH_USER")
var authPassword = os.Getenv("AUTH_PASSWORD")

func main() {
	r := gin.Default()

	if authUser == "" || authPassword == "" {
		log.Fatalf("Set AUTH_USER and AUTH_PASSWORD")
	}

	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		authUser: authPassword,
	}))
	router.Route(authorized)

	r.Run()
}
