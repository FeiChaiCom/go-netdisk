package permission

import (
	"errors"
	cfg "github.com/gaomugong/go-netdisk/config"
	"github.com/gaomugong/go-netdisk/models/db"
	R "github.com/gaomugong/go-netdisk/render"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// curl http://localhost:5000/api/permission/self_permissions/
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
		R.Error(c, err)
	}

	R.Ok(c, p)
}

// curl http://localhost:5000/api/permission/get_my_project/
func MyProjectHandler(c *gin.Context) {
	var perm *db.Permission
	var project *db.Project

	username := c.GetString("username")
	perm, err := db.GetPermissionByUsername(username)

	if err != nil {
		// No permission item found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			R.Ok(c, gin.H{})
			return
		}
		R.Error(c, err)
	}

	if perm.Role == db.ADMINISTRATOR {
		err = cfg.DB.Order("-created_at").First(&project).Error
	} else {
		err = cfg.DB.First(&project, "uuid = ?", perm.ProjectUUID).Error
	}

	if !errors.Is(err, nil) {
		R.Error(c, err)
		return
	}

	R.Ok(c, project)
}
