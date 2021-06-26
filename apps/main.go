package apps

import (
	"github.com/gaomugong/go-netdisk/apps/monitor"
	"github.com/gaomugong/go-netdisk/apps/user"
	cfg "github.com/gaomugong/go-netdisk/config"
	"github.com/gin-gonic/gin"
)

type Register func(rg *gin.RouterGroup)

var registers = []Register{
	monitor.RegisterMonitorGroup,
	user.RegisterUserGroup,
}

func InitApiRouter() *gin.Engine {

	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(cfg.ApiLogger)

	//engine := gin.Default()
	apiGroup := engine.Group("/api")
	for _, register := range registers {
		register(apiGroup)
	}

	return engine
}
