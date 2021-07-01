package matter

import (
	"github.com/gaomugong/go-netdisk/middleware"
	"github.com/gin-gonic/gin"
)

// Add matter apis to api group
func RegisterMatterGroup(rg *gin.RouterGroup) {
	users := rg.Group("/matter/")
	users.Use(middleware.JWTLoginRequired())

	users.GET("page/", PageHandler)
	users.POST("test_upload/", testUploadFile)
	users.GET("test_get_file/", testGetFile)

}
