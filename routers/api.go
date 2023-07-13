package routers

import (
	"github.com/gin-gonic/gin"
	"go_test/web"
)

func initApi(r *gin.Engine) {
	//api := r.Group("/api", gin.Logger())
	api := r.Group("/api")
	v1 := api.Group("/v1")
	v1.GET("/ping", web.Ping)
	v1.POST("/login", web.Login)
	v1.POST("/register", web.Register)
}
