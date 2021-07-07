package user

import (
	"github.com/gin-gonic/gin"
	"go-netdisk/middleware"
)

// Add user apis to api group
func RegisterUserGroup(rg *gin.RouterGroup) {
	// users := rg.Group("/user/").Use(middleware.JWTLoginRequired())
	users := rg.Group("/user/").Use(middleware.LoginRequired)
	{
		users.GET("me/", Me)
		users.GET("page/", PageHandler)
	}
}
