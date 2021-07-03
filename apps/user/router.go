package user

import (
	"github.com/gaomugong/go-netdisk/middleware"
	"github.com/gin-gonic/gin"
)

// Add user apis to api group
func RegisterUserGroup(rg *gin.RouterGroup) {
	users := rg.Group("/user/").Use(middleware.JWTLoginRequired())
	{
		users.GET("page/", PageHandler)
	}
}
