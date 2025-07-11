package main

import (
	"github.com/gin-gonic/gin"
	"take-out/global"
	"take-out/initialize"
)

func main() {
	// 初始化设置
	r := initialize.GlobalInit()

	// 设置运行环境
	gin.SetMode(global.Config.Server.Level)

	r.Run(":8080")
}
