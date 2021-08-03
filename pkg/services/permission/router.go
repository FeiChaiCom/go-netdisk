package permission

import (
	"github.com/gin-gonic/gin"
	"go-netdisk/pkg/middleware"
)

// Add permission apis to api group
func RegisterPermissionGroup(rg *gin.RouterGroup) {
	permissions := rg.Group("/permission/").Use(middleware.JWTLoginRequired())
	// permissions := rg.Group("/permission/").Use(middleware.LoginRequired)
	{
		permissions.GET("self_permissions/", SelfPermissionsHandler)
		permissions.GET("get_my_project/", MyProjectHandler)
	}
}
