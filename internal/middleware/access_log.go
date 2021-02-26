package middleware

import (
	"bytes"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/lughong/blog-service/global"
	"github.com/lughong/blog-service/pkg/logger"
)

type AccessLogWrite struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w AccessLogWrite) Write(p []byte) (int, error) {
	if _, err := w.body.Write(p); err != nil {
		return 0, err
	}

	return w.ResponseWriter.Write(p)
}

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyWrite := AccessLogWrite{
			body:           bytes.NewBuffer([]byte("")),
			ResponseWriter: c.Writer,
		}
		c.Writer = bodyWrite

		beginTime := time.Now().Unix()
		c.Next()
		endTime := time.Now().Unix()

		files := logger.Files{
			"request":  c.Request.PostForm.Encode(),
			"response": bodyWrite.body.String(),
		}
		global.Logger.WithFiles(files).Infof(
			"access log: method: %s, status_code: %d, begin_time: %d, end_time: %d",
			c.Request.Method,
			bodyWrite.Status(),
			beginTime,
			endTime,
		)
	}
}
