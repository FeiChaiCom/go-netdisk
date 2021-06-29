package common

import (
	"github.com/gin-gonic/gin"
)

func LoginRequiredMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// log.Println("before next" + c.FullPath())
		c.Next()
		// log.Println("after next" + c.FullPath())
	}
}
