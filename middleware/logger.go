package middleware

import (
	"Starry/config"
	"fmt"
	"os"
	"path"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

/*
 *
 */
func Logger() gin.HandlerFunc {
	// 日志文件路径
	logFile := path.Join(config.LogFilePath, config.LogFileName)
	// 写入文件
	src, err := os.OpenFile(logFile, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("日志错误：写入失败！", err)
	}

	// 实例化
	logger := logrus.New()
	// 设置日志级别
	logger.SetLevel(logrus.DebugLevel)
	// 设置输出
	logger.Out = src

}
