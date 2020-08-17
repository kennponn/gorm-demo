package main

import (
	"github.com/gin-gonic/gin"
	"go-demo/controller"
)

func main() {
	r := gin.Default()
	r.POST("/api/auth/register", controller.Register)
	r.Run(":8070")
}
