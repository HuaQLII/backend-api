package login

import (
	"github.com/gin-gonic/gin"
	"metadata-api/global"
	"metadata-api/modles"
	"net/http"
)

// UserController 将抽离的出来的接口返回值
// 放到这里并挂载到UserController结构体上,实现接口的继承
// 那么在接口那边方法引用时要先实例化这个结构体，然后再调用方法
// （查看的话请跳转到userRouter看具体使用）
type LoginController struct {
	// 这里可以放一些公共的方法比如自定状态码等
}

// 用户登录
func (LC LoginController) RegisterHandler(c *gin.Context) {
	var user modles.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 检查用户名是否已存在
	var count int64
	global.DB.Model(&modles.User{}).Where("username = ?", user.Username).Count(&count)
	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username already exists"})
		return
	}

	// 创建用户
	if err := global.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user created"})
}

// 登录处理函数
func (LC LoginController) LoginHandler(c *gin.Context) {
	var user modles.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 检查用户名和密码是否匹配
	var count int64
	global.DB.Model(&modles.User{}).Where("username = ? and password = ?", user.Username, user.Password).Count(&count)
	if count == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "login success"})
}
