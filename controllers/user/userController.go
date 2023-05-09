package user

import (
	"github.com/gin-gonic/gin"
	"metadata-api/controllers"
	"metadata-api/global"
	"metadata-api/modles"
)

// UserController 将抽离的出来的接口返回值
// 放到这里并挂载到UserController结构体上,实现接口的继承
// 那么在接口那边方法引用时要先实例化这个结构体，然后再调用方法
// （查看的话请跳转到userRouter看具体使用）
type UserController struct {
	// 这里可以放一些公共的方法比如自定状态码等
	controllers.CodeController
}

// Index 用户列表
func (UC UserController) Index(context *gin.Context) {
	users := []modles.User{}
	global.DB.Find(&users)
	context.JSON(200, gin.H{
		"data": users,
	})
	//UC.Success(context)这里可以使用继承的CodeController方法
}
func (UC UserController) LiInfo(context *gin.Context) {
	//var user modles.User       // 这里要注意，如果是单条数据的话，要用结构体，如果是多条数据的话，要用切片
	liInfos := []modles.User{} // 这里要注意，如果是单条数据的话，要用结构体，如果是多条数据的话，要用切片
	global.DB.Where("username like ?", "li%").Find(&liInfos)
	context.JSON(200, gin.H{
		"data": liInfos,
	})
	//UC.Success(context)这里可以使用继承的CodeController方法
}

// Add 添加新的用户

func (UC UserController) Add(context *gin.Context) {
	user := modles.User{
		Username: "zhang",
		Password: "12346",
	}
	global.DB.Create(&user) // 通过数据的指针来创建
	context.JSON(200, user)

}

// Edit 编辑用户
func (UC UserController) Edit(context *gin.Context) {
	context.JSON(200, gin.H{"message": "这里是编辑用户"})
}

// Delete 删除用户
func (UC UserController) Delete(context *gin.Context) {
	context.JSON(200, gin.H{"message": "这里是删除用户用户"})
}
