package router

import (
	"Starry/config"

	"github.com/gin-gonic/gin"
)

/*
 * gin路由启动函数
 */
func Run() {
	// 判断运行模式
	if config.RunMode == "release" {
		gin.SetMode(gin.ReleaseMode)
		gin.DisableConsoleColor()
	}

	// 创建gin engine
	r := gin.New()

	// 全局中间件
	// 使用Logger中间件
	if config.RunMode == "release" {
		r.Use(gin.Logger())
	} else {
		r.Use(gin.Logger())
	}
}
