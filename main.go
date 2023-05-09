package main

import (
	"github.com/gin-gonic/gin"
	"metadata-api/global"
	"metadata-api/routers"
)

func main() {
	//路由初始化
	router := gin.Default()
	routers.ApiAllRouterInit(router)
	//数据库初始化
	global.GormInit()
	router.Run(":8888")

}
