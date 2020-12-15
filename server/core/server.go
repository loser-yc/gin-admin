package core

import (
	"github.com/gin-gonic/gin"
	"go-admin/router"
)

//启动服务
func StartServer()  {
	r := gin.Default()
	initRouter(r)
	r.Run(":8000")
}


//初始化路由
func initRouter(r *gin.Engine)  {
	PrivateGroup := r.Group("")
	router.InitUserRouter(PrivateGroup)
}
