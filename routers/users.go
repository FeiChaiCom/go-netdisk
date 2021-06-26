package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Add user apis to api group
func registerUserGroup(rg *gin.RouterGroup) {
	monitors := rg.Group("/users/")
	monitors.GET("", userList)
}

func userList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": true,
		"data":   []string{"zhangsan", "lisi"},
	})
}
