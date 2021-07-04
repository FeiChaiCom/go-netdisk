package preference

import (
	"github.com/gin-gonic/gin"
	cfg "go-netdisk/config"
	"go-netdisk/models/db"
	R "go-netdisk/render"
)

// curl -X POST http://localhost:5000/api/preference/fetch/
func FetchHandler(c *gin.Context) {
	var p db.Preference
	if err := cfg.DB.First(&p).Error; err != nil {
		R.Error(c, err)
	}
	R.Ok(c, p)
}
