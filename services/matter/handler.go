package matter

import (
	"fmt"
	"github.com/gin-gonic/gin"
	cfg "go-netdisk/config"
	"go-netdisk/models/db"
	"go-netdisk/models/form"
	R "go-netdisk/render"
	"log"
	"os"
	"strings"
)

// Get matter list with pagination
// curl http://localhost:5000/api/matter/page/?page=1&pageSize=20&orderCreateTime=DESC&puuid=root&orderDir=DESC
func PageHandler(c *gin.Context) {
	var p form.PageParam
	if err := c.ShouldBindQuery(&p); err != nil {
		R.Error(c, err)
		return
	}

	username := c.GetString("username")
	matters, totalItems, totalPages := db.GetAllMatters(username, p.PUUID, p.Name, p.Page, p.PageSize, p.OrderCreateTime)
	R.Ok(c, gin.H{
		"totalPage":  totalPages,
		"totalItems": totalItems,
		"data":       matters,
	})
}

// Delete matter file or directory
func DeleteMatterHandler(c *gin.Context) {
	var p form.BaseQueryParam
	if err := c.ShouldBind(&p); err != nil {
		R.Error(c, err)
		return
	}

	if err := db.DeleteMatterByUUID(p.UUID); err != nil {
		R.Error(c, err)
		return
	}

	R.Ok(c, p.UUID)
}

// Get matter detail info
// curl http://localhost:5000/api/matter/get_detail/?uuid=5cfa8798-fe3e-4ffa-a0ba-b9afd88003f5
func DetailHandler(c *gin.Context) {
	var p form.BaseQueryParam
	if err := c.ShouldBind(&p); err != nil {
		R.Error(c, err)
		return
	}

	matter, err := db.GetMatterByUUID(p.UUID)
	if err != nil {
		R.Error(c, err)
		return
	}

	// Add parent info
	if matter.PUUID != cfg.MatterRootUUID {
		parent, _ := db.GetMatterByUUID(matter.PUUID)
		R.Ok(c, form.SubDirDetailMatter{Matter: matter, Parent: parent})
		return
	}

	R.Ok(c, form.RootDirDetailMatter{Matter: matter, Parent: nil})
}

// Upload file to media dir
// curl -X POST http://localhost:5000/api/matter/upload/ \
//  -F "file=@/tmp/log.tar.gz" \
//  -H "Content-Type: multipart/form-data"
func UploadFileHandler(c *gin.Context) {
	var p form.UploadParam
	if err := c.ShouldBind(&p); err != nil {
		R.Error(c, err)
		return
	}

	// Save file to local dir
	parentDir := ""
	if p.PUUID != cfg.MatterRootUUID {
		if pDir, err := db.GetMatterByUUID(p.PUUID); err == nil {
			parentDir = pDir.Path
		}
	}

	filePath := parentDir + "/" + p.File.Filename
	realFilePath := strings.Join([]string{cfg.MatterRoot, filePath}, "/")
	if err := c.SaveUploadedFile(p.File, realFilePath); err != nil {
		R.Error(c, err)
		return
	}

	username := c.GetString("username")
	matter, err := db.CreateMatter(username, p.UserUUID, p.PUUID, filePath, p.File)
	if err != nil {
		R.Error(c, err)
		return
	}

	R.OkWithMsg(c, matter, fmt.Sprintf("upload <%s> success", p.File.Filename))
}

// Download matter file as attachment
// curl http://localhost:5000/api/matter/download/?name=log.tar.gz
func DownloadFileHandler(c *gin.Context) {
	name := c.DefaultQuery("name", "")
	matterUUID := c.Param("uuid")

	matter, err := db.GetMatterByUUID(matterUUID)
	if err != nil {
		R.Error(c, err)
		return
	}

	if name == "" {
		name = matter.Name
	}

	realPath := cfg.MatterRoot + matter.File
	c.FileAttachment(realPath, name)
	// c.File(matter.Path)

	// Increment download times
	matter.Times++
	cfg.DB.Save(matter)
}

// Create matter dir
func CreateDirectoryHandler(c *gin.Context) {
	var p form.CreateDirParam
	if err := c.ShouldBind(&p); err != nil {
		R.Error(c, err)
	}

	// Create dir in filesystem
	parentDir := ""
	if p.PUUID != cfg.MatterRootUUID {
		if pDir, err := db.GetMatterByUUID(p.PUUID); err == nil {
			parentDir = pDir.Path
		}
	}

	path := parentDir + "/" + p.Name
	realPath := cfg.MatterRoot + path
	if err := os.MkdirAll(realPath, 0755); err != nil {
		log.Printf("mkdir <%s> error: %s", realPath, err)
	}

	username := c.GetString("username")
	matterDir, err := db.CreateDirectory(username, p.UserUUID, p.PUUID, path, p.Name)
	if err != nil {
		R.Error(c, err)
		return
	}

	R.OkWithMsg(c, matterDir, fmt.Sprintf("create dir <%s> success", matterDir.Name))
}
