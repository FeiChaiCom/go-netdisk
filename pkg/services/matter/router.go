package matter

import (
	"github.com/gin-gonic/gin"
	"go-netdisk/pkg/middlewares"
)

// Add matter apis to api group
func RegisterMatterGroup(rg *gin.RouterGroup) {
	users := rg.Group("/matter/").Use(middlewares.JWTLoginRequired())
	// users := rg.Group("/matter/").Use(middlewares.LoginRequired)
	{
		users.GET("page/", PageHandler)
		users.GET("get_detail/", DetailHandler)
		users.POST("delete/", DeleteMatterHandler)
		users.POST("upload/", UploadFileHandler)
		users.POST("create_directory/", CreateDirectoryHandler)
		users.GET("/:uuid/download/", DownloadFileHandler)
	}
}
