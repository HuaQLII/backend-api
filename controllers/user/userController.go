package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// UserController 将抽离的出来的接口返回值
// 放到这里并挂载到UserController结构体上,实现接口的继承
// 那么在接口那边方法引用时要先实例化这个结构体，然后再调用方法
// （查看的话请跳转到userRouter看具体使用）
type UserController struct {
}

// Index 用户列表
func (UC UserController) Index(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "这里是用户列表"})
}

// Add 添加新的用户
func (UC UserController) Add(context *gin.Context) {
	context.JSON(200, gin.H{"message": "这里是添加新的用户"})
}

// Edit 编辑用户
func (UC UserController) Edit(context *gin.Context) {
	context.JSON(200, gin.H{"message": "这里是编辑用户"})
}

// Delete 删除用户
func (UC UserController) Delete(context *gin.Context) {
	context.JSON(200, gin.H{"message": "这里是删除用户用户"})
}
