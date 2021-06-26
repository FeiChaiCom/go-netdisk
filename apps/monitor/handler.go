package monitor

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// For readiness probe
func healthzHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result":  true,
		"message": "ok",
	})
}
