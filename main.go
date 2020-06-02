package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tynmarket/table-viewer/router"
)

var port = os.Getenv("SERVER_PORT")
var authUser = os.Getenv("AUTH_USER")
var authPassword = os.Getenv("AUTH_PASSWORD")

func main() {
	r := gin.Default()

	if port == "" {
		port = ":8080"
	} else {
		port = ":" + port
	}

	if authUser == "" || authPassword == "" {
		log.Fatalf("Set AUTH_USER and AUTH_PASSWORD")
	}

	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		authUser: authPassword,
	}))
	router.Route(authorized)

	r.Run(port)
}
