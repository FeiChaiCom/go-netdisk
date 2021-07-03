package matter

import (
	"github.com/gaomugong/go-netdisk/middleware"
	"github.com/gin-gonic/gin"
)

// Add matter apis to api group
func RegisterMatterGroup(rg *gin.RouterGroup) {
	users := rg.Group("/matter/").Use(middleware.JWTLoginRequired())
	{
		users.GET("page/", PageHandler)
		users.GET("get_detail/", DetailHandler)
		users.POST("delete/", DeleteMatterHandler)
		users.POST("upload/", UploadFileHandler)
		users.POST("create_directory/", CreateDirectoryHandler)
		users.GET("/:uuid/download/", DownloadFileHandler)
	}
}
