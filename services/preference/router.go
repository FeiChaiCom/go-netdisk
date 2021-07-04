package preference

import (
	"github.com/gin-gonic/gin"
)

// Add preference apis to api group
func RegisterPreferenceGroup(rg *gin.RouterGroup) {
	users := rg.Group("/preference/")
	{
		users.POST("fetch/", FetchHandler)
	}
}
