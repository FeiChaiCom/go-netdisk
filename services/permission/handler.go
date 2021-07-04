package permission

import (
	"errors"
	cfg "github.com/gaomugong/go-netdisk/config"
	"github.com/gaomugong/go-netdisk/models/db"
	R "github.com/gaomugong/go-netdisk/render"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// curl -X POST http://localhost:5000/api/permission/self_permissions/
func SelfPermissionsHandler(c *gin.Context) {
	var p *db.Permission

	username := c.GetString("username")
	p, err := db.GetPermissionByUsername(username)

	if err != nil {
		// No permission item found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			R.Ok(c, gin.H{})
			return
		}
		R.FailWithError(c, err)
	}

	R.Ok(c, p)
}

// curl -X POST http://localhost:5000/api/permission/self_permissions/
func MyProjectHandler(c *gin.Context) {
	var p db.Preference
	if err := cfg.DB.First(&p).Error; err != nil {
		R.FailWithError(c, err)
	}
	R.Ok(c, p)
}
