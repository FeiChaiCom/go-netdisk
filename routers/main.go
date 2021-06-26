package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	apiGroup := r.Group("/api")

	// Register api of monitor
	{
		registerMonitorGroup(apiGroup)
	}

	// Register api of user
	{
		registerUserGroup(apiGroup)
	}

	return r
}

func healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result":  true,
		"message": "ok",
	})
}
