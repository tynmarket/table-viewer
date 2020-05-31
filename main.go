package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tynmarket/table-viewer/router"
)

func main() {
	r := gin.Default()
	r = router.Route(r)
	r.Run()
}
