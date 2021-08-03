package preference

import (
	"github.com/gin-gonic/gin"
	"go-netdisk/pkg/utils"

	"go-netdisk/pkg/db"
	"go-netdisk/pkg/db/models"
)

// curl -X POST http://localhost:5000/api/preference/fetch/
func FetchHandler(c *gin.Context) {
	var p models.Preference
	if err := db.DB.First(&p).Error; err != nil {
		utils.Error(c, err)
	}
	utils.Ok(c, p)
}
