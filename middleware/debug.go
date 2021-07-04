package middleware

import (
	"bytes"
	"github.com/gaomugong/go-netdisk/utils"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"log"
)

func RequestDebugLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		var buf bytes.Buffer

		tee := io.TeeReader(c.Request.Body, &buf)
		body, _ := ioutil.ReadAll(tee)
		c.Request.Body = ioutil.NopCloser(&buf)

		log.Println(string(body))
		log.Println(utils.PrettyJson(c.Request.Header))

		c.Next()
	}
}
