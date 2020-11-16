package middleware

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//refer to https://mp.weixin.qq.com/s/gBWEHe20Lv_2wBSlM2WeVA

func Valid(c *gin.Context) {
	return
}

func LoggerToFile() gin.HandlerFunc {
	fileName := "ginlog.log"
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(err)
	}

	// 实例化
	Logger := logrus.New()

	// 设置输出
	Logger.Out = f

	// 设置log级别
	Logger.SetLevel(logrus.DebugLevel)

	// 设置log格式
	Logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	return func(c *gin.Context) {
		startTime := time.Now()

		// 处理请求操作
		c.Next()

		endTime := time.Now()

		latencyTime := endTime.Sub(startTime)

		reqMethod := c.Request.Method
		reqUri := c.Request.RequestURI
		statusCode := c.Writer.Status()
		clientIp := c.ClientIP()
		Logger.Infof("| %3d | %13v | %15s | %s | %s |", statusCode, latencyTime, clientIp, reqMethod, reqUri)
	}
}
