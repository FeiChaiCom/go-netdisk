package matter

import (
	"fmt"
	cfg "github.com/gaomugong/go-netdisk/config"
	"github.com/gaomugong/go-netdisk/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

type pageParam struct {
	Puuid           string `form:"puuid" binding:"required"`
	Name            string `form:"name"`
	Page            int    `form:"page"`
	PageSize        int    `form:"pageSize"`
	OrderCreateTime string `form:"orderCreateTime"`
}

// curl http://localhost:5000/api/matter/page/?page=1&pageSize=20&orderCreateTime=DESC&puuid=root&orderDir=DESC
func PageHandler(c *gin.Context) {
	var p pageParam
	if err := c.ShouldBindQuery(&p); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":  false,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	matters, totalItems, totalPages := models.GetAllMatters(p.Puuid, p.Name, p.Page, p.PageSize, p.OrderCreateTime)
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

// curl -X POST http://localhost:5000/api/tests/test_upload/ \
//  -F "file=@/tmp/log.tar.gz" \
//  -H "Content-Type: multipart/form-data"
func testUploadFile(c *gin.Context) {
	file, _ := c.FormFile("file")
	log.Println(file.Filename)

	dstFile := strings.Join([]string{cfg.MediaDir, file.Filename}, "/")
	log.Println(dstFile)

	if err := c.SaveUploadedFile(file, dstFile); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":  true,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result":  true,
		"message": fmt.Sprintf("upload <%s> success", file.Filename),
	})
}

// curl http://localhost:5000/api/tests/test_get_file/
func testGetFile(c *gin.Context) {
	c.FileAttachment(cfg.MediaDir, "log.tar.gz")
	// curl http://localhost:5000/api/tests/test_get_file/ -o log.tar.gz
	// c.File(cfg.MediaDir + "/log.tar.gz")
}
