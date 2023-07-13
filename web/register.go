package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type register struct {
	UserName string `form:"user_name" binding:"required"`
	Pwd      string `form:"pwd" binding:"required"`
	Phone    string `form:"phone" binding:"required,e164"`   //电话格式
	Email    string `form:"email" binding:"omitempty,email"` //为空忽略
}

func Register(c *gin.Context) {
	req := register{}
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, req)
	//c.JSON(http.StatusOK, gin.H{
	//	"message": "register",
	//})
}
