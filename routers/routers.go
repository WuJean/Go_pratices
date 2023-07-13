package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRouters(r *gin.Engine) {
	//r.Use(middleware.AuthCheck,middleware.TokenCheck)
	//r.Use(gin.Logger())
	initApi(r)
	InitCourse(r)
}
