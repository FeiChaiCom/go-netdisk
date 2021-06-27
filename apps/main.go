package apps

import (
	"github.com/gaomugong/go-netdisk/apps/demo"
	"github.com/gaomugong/go-netdisk/apps/monitor"
	"github.com/gaomugong/go-netdisk/common"
	cfg "github.com/gaomugong/go-netdisk/config"
	"github.com/gin-gonic/gin"
)

type Register func(rg *gin.RouterGroup)

var registers = []Register{
	monitor.RegisterMonitorGroup,
	demo.RegisterTestGroup,
}

func InitApiRouter() *gin.Engine {

	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(cfg.ApiLogger)
	engine.Use(common.LoginRequiredMiddleware())

	// Set a lower memory limit for multipart forms (default 32M)
	engine.MaxMultipartMemory = 8 << 20 // 8MiB

	//engine := gin.Default()
	apiGroup := engine.Group("/api")
	for _, register := range registers {
		register(apiGroup)
	}

	return engine
}
