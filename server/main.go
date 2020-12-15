package main

import (
	"go-admin/core"
)

func main()  {

	core.SetConfig() //加载配置

	core.Gorm()		//初始化数据库

	core.StartServer()	//启动服务
}
