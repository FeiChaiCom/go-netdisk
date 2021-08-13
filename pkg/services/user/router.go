package user

import (
	"github.com/gin-gonic/gin"
	"go-netdisk/pkg/middlewares"
)

// Add user apis to api group
func RegisterUserGroup(rg *gin.RouterGroup) {
	users := rg.Group("/user/").Use(middlewares.JWTLoginRequired())
	// users := rg.Group("/user/").Use(middlewares.LoginRequired)
	{
		users.GET("me/", Me)
		users.GET("page/", PageHandler)
	}
}
