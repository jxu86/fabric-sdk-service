package gintool

import (
	"bytes"
	"gas-fabric-service/common/log"
	"github.com/gin-gonic/gin"
	"time"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func Logger() gin.HandlerFunc {
	logClient := log.GetLogger("ginlog", log.INFO)

	return func(c *gin.Context) {
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		//开始时间
		start := time.Now()
		//处理请求
		c.Next()
		//结束时间
		end := time.Now()
		//执行时间
		latency := end.Sub(start)
		//path
		path := c.Request.URL.Path
		//ip
		clientIP := c.ClientIP()
		//方法
		method := c.Request.Method
		//状态
		statusCode := c.Writer.Status()
		logClient.Infof("| %3d | %13v | %15s | %s  %s |",
			statusCode,
			latency,
			clientIP,
			method,
			path,
		)
		
		logClient.Infof("Response body: ", blw.body.String())
	}
}
