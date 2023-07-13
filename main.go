package main

import (
	"github.com/gin-gonic/gin"
	"go_test/routers"
)

func main() {
	//r := gin.Default()
	r := gin.New()
	r.Use(gin.Recovery(), gin.Logger())
	routers.InitRouters(r)
	r.Run(":8080")
}
