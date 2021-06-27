package config

import (
	"fmt"
	"github.com/gaomugong/go-netdisk/common"
	"github.com/gin-gonic/gin"
)

var ApiLogger = gin.LoggerWithFormatter(func(p gin.LogFormatterParams) string {
	return fmt.Sprintf("[%s] %d | %s | %s | %d | %s\t| %s\n",
		p.TimeStamp.Format(common.SimpleTime),
		p.StatusCode,
		p.Method,
		p.Latency,
		p.BodySize,
		p.Path,
		p.ClientIP,
	)
})
