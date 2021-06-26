package apps

import (
	"github.com/gaomugong/go-netdisk/apps/monitor"
	"github.com/gaomugong/go-netdisk/apps/user"
	"github.com/gin-gonic/gin"
)

type Register func(rg *gin.RouterGroup)

var registers = []Register{
	monitor.RegisterMonitorGroup,
	user.RegisterUserGroup,
}

func InitApiRouter() *gin.Engine {
	r := gin.Default()
	apiGroup := r.Group("/api")

	for _, register := range registers {
		register(apiGroup)
	}

	return r
}
