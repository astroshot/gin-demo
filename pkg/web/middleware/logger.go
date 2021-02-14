package middleware

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"gin-demo/pkg/config"
	"gin-demo/pkg/util"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// GetLogger returns logger based on logrus
func GetLogger() gin.HandlerFunc {
	logger := config.GetLogger()

	return func(c *gin.Context) {
		// Process requestBody
		data, err := c.GetRawData()
		if err != nil {
			fmt.Println(err.Error())
		}
		dataStr := fmt.Sprintf("%v", string(data))
		// fmt.Printf("data: %v\n", string(data))
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data)) // 关键点

		// Process responseBody
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		start := time.Now()
		c.Next()
		end := time.Now()
		latency := end.Sub(start)
		latencyStr := fmt.Sprintf("%fms", float64(latency)/float64(time.Millisecond))

		// path := c.Request.URL.Path
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		var traceIDStr string
		if traceID, exist := c.Get(util.TraceIDKey); !exist {
			traceIDStr = ""
		} else {
			traceIDStr = fmt.Sprintf("%v", traceID)
		}

		logger.WithFields(logrus.Fields{
			"proto":                c.Request.Proto,
			"host":                 c.Request.Host,
			"status":               statusCode,
			"method":               method,
			"requestContentLength": c.Request.ContentLength,
			"requestHeader":        c.Request.Header,
			"trailer":              c.Request.Trailer,
			"requestBody":          dataStr,
			"URI":                  c.Request.URL.Path,
			"responseBody":         blw.body.String(),
			"clientIP":             clientIP,
			"remoteAddr":           c.Request.RemoteAddr,
			"cost":                 latencyStr,
			"traceID":              traceIDStr,
		}).Infof("")
	}
}
