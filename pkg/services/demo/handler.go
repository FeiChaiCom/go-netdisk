package demo

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go-netdisk/pkg/settings"

	"log"
	"net/http"
	"strings"
	"time"
)

type User struct {
	Name string `form:"name" json:"name" binding:"required"`
	Age  int    `form:"age" json:"age" binding:"required,gt=10"`
	// time_format has issue when set content-type to application/json, but work with form style
	BirthDay time.Time `form:"birthday" json:"birthday" time_format:"2006-01-02 15:04:05"`
}

// TODO: parsing time "2021-05-27 15:04:05" as "2006-01-02T15:04:05Z07:00": cannot parse " 15:04:05" as "T"
var birthDayValidator validator.Func = func(fl validator.FieldLevel) bool {
	birthDay, ok := fl.Field().Interface().(time.Time)
	log.Println(birthDay, ok)
	if ok {
		today := time.Now()
		return !today.Before(birthDay)
	}
	return false
}

// curl http://localhost:5000/api/tests/test_post/ -X POST -d '{"name": "miya", "age": 11}' -H 'Content-Type: application/json' -vvv
// curl http://localhost:5000/api/tests/test_post/ -X POST -F name=miya -F age=20 -F birthday="2020-11-12 11:11:11"
// curl http://localhost:5000/api/tests/test_post/ -X POST -d '{"name": "miya", "age": 1, "birthday": "2020-01-02T11:12:13.123433"}' -H 'Content-Type: application/json' -vvv
func testPostUserData(c *gin.Context) {
	var user User

	// if err := c.ShouldBindJSON(&demo); err != nil {
	// ShouldBind will guess which content-type is the request and call the specific bind type, eg. bindJSON
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result":  false,
			"message": err.Error(),
		})
		return
	}

	userJSON, _ := json.Marshal(user)
	log.Println(string(userJSON))
	c.JSON(http.StatusOK, gin.H{
		"result": true,
		"data":   user,
	})
}

// curl http://localhost:5000/api/tests/test_post_form/ -X POST -F name=miya -F age=12
func testPostUser(c *gin.Context) {
	name := c.DefaultPostForm("name", "pitou")
	age := c.PostForm("age")
	c.JSON(http.StatusOK, gin.H{
		"result": true,
		"data": map[string]interface{}{
			"name": name,
			"age":  age,
		},
	})
}

// curl http://localhost:5000/api/tests/test_get/1/eat/?age=12
func testGetUser(c *gin.Context) {
	name := c.DefaultQuery("name", "pitou")
	age := c.Query("age")
	uid := c.Param("uid")
	action := c.Param("action")
	action = strings.Trim(action, "/")

	cc := c.Copy()

	go func() {
		time.Sleep(3 * time.Second)
		log.Println("run in go func: " + cc.Request.URL.String())
	}()

	c.JSON(http.StatusOK, gin.H{
		"result": true,
		"data": gin.H{
			"name":   name,
			"age":    age,
			"uid":    uid,
			"action": action,
			"list":   []string{"zhangsan", "lisi"},
		},
	})
}

// curl http://localhost:5000/api/tests/test_redirect/
func testRedirect(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
}

//curl -X POST http://localhost:5000/api/tests/test_upload/ \
//  -F "file=@/tmp/log.tar.gz" \
//  -H "Content-Type: multipart/form-data"
func testUploadFile(c *gin.Context) {
	file, _ := c.FormFile("file")
	log.Println(file.Filename)

	dstFile := strings.Join([]string{settings.ENV.MediaDir, file.Filename}, "/")
	log.Println(dstFile)

	if err := c.SaveUploadedFile(file, dstFile); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":  false,
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
	c.FileAttachment(settings.ENV.MediaDir, "log.tar.gz")
	// curl http://localhost:5000/api/tests/test_get_file/ -o log.tar.gz
	// c.File(cfg.MediaDir + "/log.tar.gz")
}
