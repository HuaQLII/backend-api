package main

import (
	"github.com/gin-gonic/gin"
	"metadata-api/routers"
)

func main() {
	router := gin.Default()
	routers.ApiAllRouterInit(router)
	router.Run(":0113")

}
