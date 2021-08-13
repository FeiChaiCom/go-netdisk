package matter

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-netdisk/pkg/db"
	"go-netdisk/pkg/settings"
	"go-netdisk/pkg/utils"

	"go-netdisk/pkg/db/models"
	"go-netdisk/pkg/services/form"
	"log"
	"os"
	"strings"
)

// Get matter list with pagination
// curl http://localhost:5000/api/matter/page/?page=1&pageSize=20&orderCreateTime=DESC&puuid=root&orderDir=DESC
func PageHandler(c *gin.Context) {
	var p form.PageParam
	if err := c.ShouldBindQuery(&p); err != nil {
		utils.Error(c, err)
		return
	}

	username := c.GetString("username")
	matters, totalItems, totalPages := models.GetAllMatters(username, p.PUUID, p.Name, p.Page, p.PageSize, p.OrderCreateTime)
	utils.Ok(c, gin.H{
		"totalPage":  totalPages,
		"totalItems": totalItems,
		"data":       matters,
	})
}

// Delete matter file or directory
func DeleteMatterHandler(c *gin.Context) {
	var p form.BaseQueryParam
	if err := c.ShouldBind(&p); err != nil {
		utils.Error(c, err)
		return
	}

	if err := models.DeleteMatterByUUID(p.UUID); err != nil {
		utils.Error(c, err)
		return
	}

	utils.Ok(c, p.UUID)
}

// Get matter detail info
// curl http://localhost:5000/api/matter/get_detail/?uuid=5cfa8798-fe3e-4ffa-a0ba-b9afd88003f5
func DetailHandler(c *gin.Context) {
	var p form.BaseQueryParam
	if err := c.ShouldBind(&p); err != nil {
		utils.Error(c, err)
		return
	}

	matter, err := models.GetMatterByUUID(p.UUID)
	if err != nil {
		utils.Error(c, err)
		return
	}

	// Add parent info
	if matter.PUUID != settings.MatterRootUUID {
		parent, _ := models.GetMatterByUUID(matter.PUUID)
		utils.Ok(c, form.SubDirDetailMatter{Matter: matter, Parent: parent})
		return
	}

	utils.Ok(c, form.RootDirDetailMatter{Matter: matter, Parent: nil})
}

// Upload file to media dir
// curl -X POST http://localhost:5000/api/matter/upload/ \
//  -F "file=@/tmp/log.tar.gz" \
//  -H "Content-Type: multipart/form-data"
func UploadFileHandler(c *gin.Context) {
	var p form.UploadParam
	if err := c.ShouldBind(&p); err != nil {
		utils.Error(c, err)
		return
	}

	// Save file to local dir
	parentDir := ""
	if p.PUUID != settings.MatterRootUUID {
		if pDir, err := models.GetMatterByUUID(p.PUUID); err == nil {
			parentDir = pDir.Path
		}
	}

	filePath := parentDir + "/" + p.File.Filename
	realFilePath := strings.Join([]string{settings.ENV.MatterRoot, filePath}, "/")
	if err := c.SaveUploadedFile(p.File, realFilePath); err != nil {
		utils.Error(c, err)
		return
	}

	username := c.GetString("username")
	matter, err := models.CreateMatter(username, p.UserUUID, p.PUUID, filePath, p.File)
	if err != nil {
		utils.Error(c, err)
		return
	}

	utils.OkWithMsg(c, matter, fmt.Sprintf("upload <%s> success", p.File.Filename))
}

// Download matter file as attachment
// curl http://localhost:5000/api/matter/download/?name=log.tar.gz
func DownloadFileHandler(c *gin.Context) {
	name := c.DefaultQuery("name", "")
	matterUUID := c.Param("uuid")

	matter, err := models.GetMatterByUUID(matterUUID)
	if err != nil {
		utils.Error(c, err)
		return
	}

	if name == "" {
		name = matter.Name
	}

	realPath := settings.ENV.MatterRoot + matter.File
	c.FileAttachment(realPath, name)
	// c.File(matter.Path)

	// Increment download times
	matter.Times++
	db.DB.Save(matter)
}

// Create matter dir
func CreateDirectoryHandler(c *gin.Context) {
	var p form.CreateDirParam
	if err := c.ShouldBind(&p); err != nil {
		utils.Error(c, err)
	}

	// Create dir in filesystem
	parentDir := ""
	if p.PUUID != settings.MatterRootUUID {
		if pDir, err := models.GetMatterByUUID(p.PUUID); err == nil {
			parentDir = pDir.Path
		}
	}

	path := parentDir + "/" + p.Name
	realPath := settings.ENV.MatterRoot + path
	if err := os.MkdirAll(realPath, 0755); err != nil {
		log.Printf("mkdir <%s> error: %s", realPath, err)
	}

	username := c.GetString("username")
	matterDir, err := models.CreateDirectory(username, p.UserUUID, p.PUUID, path, p.Name)
	if err != nil {
		utils.Error(c, err)
		return
	}

	utils.OkWithMsg(c, matterDir, fmt.Sprintf("create dir <%s> success", matterDir.Name))
}
