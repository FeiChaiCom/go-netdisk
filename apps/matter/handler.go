package matter

import (
	"fmt"
	cfg "github.com/gaomugong/go-netdisk/config"
	"github.com/gaomugong/go-netdisk/models/db"
	"github.com/gin-gonic/gin"
	"log"
	"mime/multipart"
	"net/http"
	"strings"
	"time"
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

// curl http://localhost:5000/api/matter/page/?page=1&pageSize=20&orderCreateTime=DESC&puuid=root&orderDir=DESC
func PageHandler(c *gin.Context) {
	var p PageParam
	if err := c.ShouldBindQuery(&p); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":  false,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	matters, totalItems, totalPages := db.GetAllMatters(p.Puuid, p.Name, p.Page, p.PageSize, p.OrderCreateTime)
	log.Printf("%#v %d %d\n", p, totalItems, totalPages)

	c.JSON(http.StatusOK, gin.H{
		"result": true,
		"data": gin.H{
			"totalPage":  totalPages,
			"totalItems": totalItems,
			"data":       matters,
		},
		"message": "ok",
	})
}

// curl http://localhost:5000/api/matter/get_detail/?uuid=5cfa8798-fe3e-4ffa-a0ba-b9afd88003f5
func DetailHandler(c *gin.Context) {
	detailParam := &BaseQueryParam{}
	if err := c.ShouldBind(detailParam); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":  false,
			"message": err.Error(),
		})
		return
	}

	matter, err := db.GetMatterByUUID(detailParam.UUID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":  false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": true,
		"data":   matter,
	})
}

// curl -X POST http://localhost:5000/api/matter/upload/ \
//  -F "file=@/tmp/log.tar.gz" \
//  -H "Content-Type: multipart/form-data"
func UploadFileHandler(c *gin.Context) {
	var uploadParam UploadParam
	log.Printf("upload context: %#v", c.Request)
	if err := c.ShouldBind(&uploadParam); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":  false,
			"message": err.Error(),
		})
		return
	}

	log.Printf("upload param: %#v", uploadParam)
	file := uploadParam.File
	dstFilePath := strings.Join([]string{cfg.MediaDir, file.Filename}, "/")
	if err := c.SaveUploadedFile(file, dstFilePath); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":  true,
			"message": err.Error(),
		})
		return
	}

	matter, err := db.CreateMatter(uploadParam.UserUUID, uploadParam.Puuid, dstFilePath, uploadParam.File)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":  true,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result":  true,
		"data":    matter,
		"message": fmt.Sprintf("upload <%s> success", file.Filename),
	})
}

// curl http://localhost:5000/api/tests/test_get_file/
func testGetFile(c *gin.Context) {
	tmpFileName := fmt.Sprintf("download_%d", time.Now().Unix())
	name := c.DefaultQuery("name", tmpFileName)
	matterUUID := c.Param("uuid")

	matter, err := db.GetMatterByUUID(matterUUID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":  false,
			"message": err.Error(),
		})
		return
	}

	log.Printf("%s, %s, %s, %s, %#v", name, matterUUID, matter.Path, matter.Name, matter)
	c.FileAttachment(matter.Path, matter.Name)
	// curl http://localhost:5000/api/tests/test_get_file/ -o log.tar.gz
	// c.File(cfg.MediaDir + "/log.tar.gz")
}

func CreateDirectoryHandler(c *gin.Context) {
	var createDirParam CreateDirParam
	if err := c.ShouldBind(&createDirParam); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":  true,
			"message": err.Error(),
		})
	}

	matterDir, err := db.CreateDirectory(createDirParam.UserUUID, createDirParam.Puuid, createDirParam.Name)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":  true,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result":  true,
		"data":    matterDir,
		"message": fmt.Sprintf("create dir <%s> success", matterDir.Name),
	})
}
