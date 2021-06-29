package matter

import (
	"github.com/gin-gonic/gin"
)

// Add matter apis to api group
func RegisterMatterGroup(rg *gin.RouterGroup) {
	users := rg.Group("/matter/")

	users.GET("page/", PageHandler)
	users.POST("test_upload/", testUploadFile)
	users.GET("test_get_file/", testGetFile)

}
