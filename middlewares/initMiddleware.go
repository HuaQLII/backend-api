package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func GoRtine(c *gin.Context) {
	cCp := c.Copy() // 这里是为了防止协程里面的c被提前释放，所以要复制一份
	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("我是协程" + cCp.Request.URL.Path)
	}()
}
