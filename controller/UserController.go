package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-demo/common"
	"go-demo/model"
)

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	telphone := c.PostForm("telphone")
	fmt.Printf("%s\t%s\t%s", username, password, telphone)
	db := common.InitDB()
	user := model.User{Username: username, Password: password, Telphone: telphone}
	db.Save(&user)

}
