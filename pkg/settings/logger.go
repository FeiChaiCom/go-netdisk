package settings

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var APILogger = gin.LoggerWithFormatter(func(p gin.LogFormatterParams) string {
	return fmt.Sprintf("[%s] %d | %s | %s | %d | %s\t| %s\n",
		p.TimeStamp.Format(SimpleTime),
		p.StatusCode,
		p.Method,
		p.Latency,
		p.BodySize,
		p.Path,
		p.ClientIP,
	)
})
