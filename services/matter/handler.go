package matter

import (
	"fmt"
	cfg "github.com/gaomugong/go-netdisk/config"
	"github.com/gaomugong/go-netdisk/models/db"
	R "github.com/gaomugong/go-netdisk/render"
	"github.com/gin-gonic/gin"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

type PageParam struct {
	Puuid           string `form:"puuid"`
	Name            string `form:"name"`
	Page            int    `form:"page"`
	PageSize        int    `form:"pageSize"`
	OrderCreateTime string `form:"orderCreateTime"`
}

type BaseQueryParam struct {
	UUID string `form:"uuid" binding:"required"`
}

type BasePostParam struct {
	UserUUID string `form:"userUuid" binding:"required"`
	Puuid    string `form:"puuid" binding:"required"`
}

type UploadParam struct {
	BasePostParam
	File *multipart.FileHeader `form:"file" binding:"required"`
}

type CreateDirParam struct {
	BasePostParam
	Name string `form:"name" binding:"required"`
}

// Get matter list with pagination
// curl http://localhost:5000/api/matter/page/?page=1&pageSize=20&orderCreateTime=DESC&puuid=root&orderDir=DESC
func PageHandler(c *gin.Context) {
	var p PageParam
	if err := c.ShouldBindQuery(&p); err != nil {
		R.FailWithError(c, err)
		return
	}

	matters, totalItems, totalPages := db.GetAllMatters(p.Puuid, p.Name, p.Page, p.PageSize, p.OrderCreateTime)
	// log.Printf("%#v %d %d\n", p, totalItems, totalPages)
	R.Ok(c, gin.H{
		"totalPage":  totalPages,
		"totalItems": totalItems,
		"data":       matters,
	})
}

// Delete matter file or directory
func DeleteMatterHandler(c *gin.Context) {
	var p BaseQueryParam
	if err := c.ShouldBind(&p); err != nil {
		R.FailWithError(c, err)
		return
	}

	if err := db.DeleteMatterByUUID(p.UUID); err != nil {
		R.FailWithError(c, err)
		return
	}

	R.Ok(c, p.UUID)
}

// Get matter detail info
// curl http://localhost:5000/api/matter/get_detail/?uuid=5cfa8798-fe3e-4ffa-a0ba-b9afd88003f5
func DetailHandler(c *gin.Context) {
	var p BaseQueryParam
	if err := c.ShouldBind(&p); err != nil {
		R.FailWithError(c, err)
		return
	}

	matter, err := db.GetMatterByUUID(p.UUID)
	if err != nil {
		R.FailWithError(c, err)
		return
	}

	// Add parent info
	if matter.PUUID != cfg.MatterRootUUID {
		parent, _ := db.GetMatterByUUID(matter.PUUID)
		R.Ok(c, db.SubDirDetailMatter{Matter: matter, Parent: parent})
		return
	}

	R.Ok(c, db.RootDirDetailMatter{Matter: matter, Parent: nil})
}

// Upload file to media dir
// curl -X POST http://localhost:5000/api/matter/upload/ \
//  -F "file=@/tmp/log.tar.gz" \
//  -H "Content-Type: multipart/form-data"
func UploadFileHandler(c *gin.Context) {
	var p UploadParam
	if err := c.ShouldBind(&p); err != nil {
		R.FailWithError(c, err)
		return
	}

	// Save file to local dir
	parentDir := ""
	if p.Puuid != cfg.MatterRootUUID {
		if pDir, err := db.GetMatterByUUID(p.Puuid); err == nil {
			parentDir = pDir.Path
		}
	}

	filePath := parentDir + "/" + p.File.Filename
	realFilePath := strings.Join([]string{cfg.MatterRoot, filePath}, "/")
	if err := c.SaveUploadedFile(p.File, realFilePath); err != nil {
		R.FailWithError(c, err)
		return
	}

	username := c.GetString("username")
	matter, err := db.CreateMatter(username, p.UserUUID, p.Puuid, filePath, p.File)
	if err != nil {
		R.FailWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result":  true,
		"data":    matter,
		"message": fmt.Sprintf("upload <%s> success", p.File.Filename),
	})
}

// Download matter file as attachment
// curl http://localhost:5000/api/matter/download/?name=log.tar.gz
func DownloadFileHandler(c *gin.Context) {
	name := c.DefaultQuery("name", "")
	matterUUID := c.Param("uuid")

	matter, err := db.GetMatterByUUID(matterUUID)
	if err != nil {
		R.FailWithError(c, err)
		return
	}

	if name == "" {
		name = matter.Name
	}

	realPath := cfg.MatterRoot + matter.File
	c.FileAttachment(realPath, name)
	// c.File(matter.Path)
}

// Create matter dir
func CreateDirectoryHandler(c *gin.Context) {
	var p CreateDirParam
	if err := c.ShouldBind(&p); err != nil {
		R.FailWithError(c, err)
	}

	// Create dir in filesystem
	parentDir := ""
	if p.Puuid != cfg.MatterRootUUID {
		if pDir, err := db.GetMatterByUUID(p.Puuid); err == nil {
			parentDir = pDir.Path
		}
	}

	path := parentDir + "/" + p.Name
	realPath := cfg.MatterRoot + path
	if err := os.MkdirAll(realPath, 0755); err != nil {
		log.Printf("mkdir <%s> error: %s", realPath, err)
	}

	username := c.GetString("username")
	matterDir, err := db.CreateDirectory(username, p.UserUUID, p.Puuid, path, p.Name)
	if err != nil {
		R.FailWithError(c, err)
		return
	}

	R.OkWithMessage(c, matterDir, fmt.Sprintf("create dir <%s> success", matterDir.Name))
}
