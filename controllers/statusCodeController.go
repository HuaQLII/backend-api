package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type CodeController struct {
}

// Success 自定义状态码并存储到CodeController结构体中
func (ConC CodeController) Success(context *gin.Context) {
	context.String(http.StatusOK, "成功")

}

// Error 自定义状态码并存储到CodeController结构体中
func (ConC CodeController) Error(context *gin.Context) {
	context.String(http.StatusOK, "失败")

}
