package routers

import (
	"github.com/gin-gonic/gin"
	"metadata-api/controllers/user"
)

func ApiAllRouterInit(r *gin.Engine) {

	// 前端GET路由组
	getFrontEnd := r.Group("/v1/user")
	{
		getFrontEnd.GET(`/index`, user.UserController{}.Index)   // 用户信息列表
		getFrontEnd.GET("/add", user.UserController{}.Add)       // 用户添加
		getFrontEnd.GET("/edit", user.UserController{}.Edit)     // 用户信息编辑
		getFrontEnd.GET("/delete", user.UserController{}.Delete) // 用户删除
	}

	// 前端POST路由组
	postBackEnd := r.Group("/v2")
	{
		postBackEnd.POST("/what", func(c *gin.Context) {
			c.JSON(200, gin.H{ // 返回JSON格式的数据
				"message": "这里是V2路由组",
			})
		})

	}
}

//func ApiAllRouterInit2(router *gin.Engine) {
//
//	// 前端GET路由组
//	getFrontEnd := router.Group("/v1")
//	{
//		getFrontEnd.GET("/what", func(c *gin.Context) { // 返回JSON格式的数据
//			c.JSON(200, gin.H{
//				"message": "这里是V1路由组",
//			})
//		})
//	}
//
//	// 前端POST路由组
//	postBackEnd := router.Group("/v2")
//	{
//		postBackEnd.POST("/what", func(c *gin.Context) {
//			c.JSON(200, gin.H{ // 返回JSON格式的数据
//				"message": "这里是V2路由组",
//			})
//		})
//
//	}
//
//	router.Run(":8080")
//}
