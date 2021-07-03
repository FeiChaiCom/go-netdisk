package monitor

import (
	"github.com/gin-gonic/gin"
)

// Add health apis to api group
func RegisterMonitorGroup(rg *gin.RouterGroup) {
	monitors := rg.Group("/monitors/")
	monitors.GET("/healthz", healthzHandler)
}
