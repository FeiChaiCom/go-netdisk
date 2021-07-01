package login

import (
	"github.com/gin-gonic/gin"
)

// Add user apis to api group
func RegisterLoginGroup(rg *gin.RouterGroup) {
	users := rg.Group("/account/")

	users.POST("login/", AuthHandler)
	users.POST("register/", RegisterHandler)

}