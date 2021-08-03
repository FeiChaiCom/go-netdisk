package preference

import (
	"github.com/gin-gonic/gin"
)

// Add preference apis to api group
func RegisterPreferenceGroup(rg *gin.RouterGroup) {
	preferences := rg.Group("/preference/")
	{
		preferences.POST("fetch/", FetchHandler)
	}
}
