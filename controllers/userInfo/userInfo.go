package userInfo

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// UserInfoController 将抽离的出来的接口返回值
// 放到这里并挂载到UserInfoController结构体上,实现接口的继承
// 那么在接口那边方法引用时要先实例化这个结构体，然后再调用方法
// （查看的话请跳转到userRouter看具体使用）
type UserInfoController struct {
}

// Add 添加新的用户
func (UC UserInfoController) Add(context *gin.Context) {
	context.JSON(200, gin.H{"message": "这里是添加用户信息"})
}

// Edit 编辑用户
func (UC UserInfoController) DoUpload(c *gin.Context) {
	// 单文件
	file, _ := c.FormFile("file")
	log.Println(file.Filename)

	dst := "./static/upload" + file.Filename
	// 上传文件至指定的完整文件路径
	c.SaveUploadedFile(file, dst)

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
