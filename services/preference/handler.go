package preference

import (
	"github.com/gin-gonic/gin"

	"go-netdisk/db"
	"go-netdisk/db/models"
	R "go-netdisk/render"
)

// curl -X POST http://localhost:5000/api/preference/fetch/
func FetchHandler(c *gin.Context) {
	var p models.Preference
	if err := db.DB.First(&p).Error; err != nil {
		R.Error(c, err)
	}
	R.Ok(c, p)
}
