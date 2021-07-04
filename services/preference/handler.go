package preference

import (
	cfg "github.com/gaomugong/go-netdisk/config"
	"github.com/gaomugong/go-netdisk/models/db"
	R "github.com/gaomugong/go-netdisk/render"
	"github.com/gin-gonic/gin"
)

// curl -X POST http://localhost:5000/api/preference/fetch/
func FetchHandler(c *gin.Context) {
	var p db.Preference
	if err := cfg.DB.First(&p).Error; err != nil {
		R.Error(c, err)
	}
	R.Ok(c, p)
}
