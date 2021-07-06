package login

import (
	"github.com/gin-gonic/gin"
)

// Add user apis to api group
func RegisterLoginGroup(rg *gin.RouterGroup) {
	users := rg.Group("/account/")

	users.POST("login/", JwtLoginHandler)
	users.POST("logout/", JwtLogoutHandler)
	users.POST("register/", RegisterHandler)
	users.GET("login_page/", LoginPageHandler)
	users.GET("login_success/", LoginSuccessHandler)

}
