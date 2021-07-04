package permission

import (
	"github.com/gaomugong/go-netdisk/middleware"
	"github.com/gin-gonic/gin"
)

// Add permission apis to api group
func RegisterPermissionGroup(rg *gin.RouterGroup) {
	permissions := rg.Group("/permission/").Use(middleware.JWTLoginRequired())
	{
		permissions.GET("self_permissions/", SelfPermissionsHandler)
		permissions.GET("get_my_project/", MyProjectHandler)
	}
}
