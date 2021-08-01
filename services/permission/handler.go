package permission

import (
	"errors"
	"github.com/gin-gonic/gin"

	"go-netdisk/db"
	"go-netdisk/db/models"
	R "go-netdisk/render"
	"gorm.io/gorm"
)

// curl http://localhost:5000/api/permission/self_permissions/
func SelfPermissionsHandler(c *gin.Context) {
	var p *models.Permission

	username := c.GetString("username")
	p, err := models.GetPermissionByUsername(username)

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
	var perm *models.Permission
	var project *models.Project

	username := c.GetString("username")
	perm, err := models.GetPermissionByUsername(username)

	if err != nil {
		// No permission item found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			R.Ok(c, gin.H{})
			return
		}
		R.Error(c, err)
	}

	if perm.Role == models.ADMINISTRATOR {
		err = db.DB.Order("-created_at").First(&project).Error
	} else {
		err = db.DB.First(&project, "uuid = ?", perm.ProjectUUID).Error
	}

	if !errors.Is(err, nil) {
		R.Error(c, err)
		return
	}

	R.Ok(c, project)
}
