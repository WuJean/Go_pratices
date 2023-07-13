package routers

import (
	"github.com/gin-gonic/gin"
	"go_test/middleware"
	"go_test/web"
)

func InitCourse(r *gin.Engine) {
	v1 := r.Group("/v1", middleware.TokenCheck, middleware.AuthCheck)
	//course := v1.Group("/course")
	v1.POST("/course", web.CreateCourse)
	v1.GET("/course", web.GetCourse)
	v1.PUT("/course", web.EditCourse)
	v1.DELETE("/course", web.DeleteCourse)

	v2 := r.Group("/v2")
	v2.POST("/course", web.CreateCourseV2)
}
