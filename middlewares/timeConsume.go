package middlewares

import (
	"github.com/gin-gonic/gin"
	"time"
)

func TimeConsume(context *gin.Context) {
	startTime := time.Now().UnixNano()
	context.Next() // 执行下一个方法，最后再执行下面的代码，匹配路由的时候会先执行这之前的代码，然后再执行下面的代码
	endTime := time.Now().UnixNano()
	costTime := (endTime - startTime) / 1000
	println("耗时：", costTime, "微秒")
}
