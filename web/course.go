package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateCourse(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "create course",
	})
}

func EditCourse(c *gin.Context) {

}

func GetCourse(c *gin.Context) {

}

func DeleteCourse(c *gin.Context) {

}

func CreateCourseV2(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "create course V2",
	})
}
