package user

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"strings"
	"time"
)

type User struct {
	Name string `form:"name" json:"name"`
	Age  int    `form:"age" json:"age" binding:"required"`
	//BirthDay time.Time `json:"birthday" binding:"required"`
}

// TODO: parsing time "2021-05-27 15:04:05" as "2006-01-02T15:04:05Z07:00": cannot parse " 15:04:05" as "T"
var birthDayValidator validator.Func = func(fl validator.FieldLevel) bool {
	birthDay, ok := fl.Field().Interface().(time.Time)
	log.Println(birthDay, ok)
	if ok {
		today := time.Now()
		if today.Before(birthDay) {
			return false
		}
		return true
	}
	return false
}

//curl  curl http://localhost:5000/api/users/test_post/ -X POST -d '{"name": "miya", "age": 12}'
func testPostUserData(c *gin.Context) {
	var user User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result":  false,
			"message": err.Error(),
		})
		return
	}

	userJson, _ := json.Marshal(user)
	log.Println(string(userJson))
	c.JSON(http.StatusOK, gin.H{
		"result": true,
		"data":   user,
	})
}

//curl http://localhost:5000/api/users/test_post_form/ -X POST -F name=miya -F age=12
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

// curl http://localhost:5000/api/users/test_get/1/eat/?age=12
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
		"data": map[string]interface{}{
			"name": name,
			"age":  age,
			"uid":  uid,
			"list": []string{"zhangsan", "lisi"},
		},
	})
}

// curl http://localhost:5000/api/users/test_redirect/
func testRedirect(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
}
