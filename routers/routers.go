package routers

import (
	"github.com/gin-gonic/gin"
	"metadata-api/controllers/login"
	"metadata-api/controllers/user"
	"metadata-api/controllers/userInfo"
	"metadata-api/middlewares"
)

func ApiAllRouterInit(r *gin.Engine) {
	r.Use(middlewares.Cors()) // 跨域中间件Use可以使用多个中间件.且是全局中间件
	// 前端GET路由组
	getFrontEnd := r.Group("/v1/user") // 路由组中间件也可以在这里使用
	{
		getFrontEnd.GET("/lisiinfo", user.UserController{}.LiInfo) // 用户登录
		getFrontEnd.GET("/index", user.UserController{}.Index)     // 用户列表
		getFrontEnd.POST("/add", user.UserController{}.Add)        // 用户添加
		getFrontEnd.GET("/edit", user.UserController{}.Edit)       // 用户信息编辑
		getFrontEnd.GET("/delete", user.UserController{}.Delete)   // 用户删除
	}

	//前端POST路由组
	postBackEnd := r.Group("/v2")
	{
		postBackEnd.POST("/doupload", userInfo.UserInfoController{}.DoUpload)  // 用户登录
		postBackEnd.POST("/register", login.LoginController{}.RegisterHandler) // 用户注册
		postBackEnd.POST("/login", login.LoginController{}.LoginHandler)       // 用户登录
	}

}
