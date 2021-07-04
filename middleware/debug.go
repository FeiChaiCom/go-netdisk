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

		// log.Printf("\n" +
		// 	"FullPath:\t%s\n" +
		// 	"Form:\t%s\n" +
		// 	"PostForm:\t%s\n",
		// 	c.FullPath(),
		// 	utils.PrettyJson(c.Request.Form),
		// 	utils.PrettyJson(c.Request.PostForm),
		// )
		log.Println(string(body))
		log.Println(utils.PrettyJson(c.Request.Header))

		c.Next()
	}
}
