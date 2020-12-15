package router

import (
	"github.com/gin-gonic/gin"
	v1 "go-admin/api/v1"
)

func InitUserRouter(Router *gin.RouterGroup)  {
	Router.GET("/user", v1.Login)
}