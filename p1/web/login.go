package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type loginReq struct {
	UserName string `form:"user_name"`
	Pwd      string `form:"pwd"`
}

func Login(c *gin.Context) {
	req := &loginReq{}
	c.Bind(&req)
	//c.BindQuery()	//GET方法
	//c.BindJSON()
	//err := c.ShouldBind()	//错误不会直接返回错误信息
	//c.JSON(http.StatusOK, gin.H{
	//	"message": "login",
	//})
	c.JSON(http.StatusOK, req)
}
