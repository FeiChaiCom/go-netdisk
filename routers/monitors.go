package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Add health apis to api group
func registerMonitorGroup(rg *gin.RouterGroup) {
	monitors := rg.Group("/monitors/")
	monitors.GET("/healthz", healthzHandler)
}

// For readiness probe
func healthzHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result":  true,
		"message": "ok",
	})
}
