package user

import (
	"github.com/gin-gonic/gin"
)

// Add user apis to api group
func RegisterUserGroup(rg *gin.RouterGroup) {
	users := rg.Group("/user/")

	users.GET("page/", PageHandler)

}
