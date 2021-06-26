package routers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// Add user apis to api group
func registerUserGroup(rg *gin.RouterGroup) {
	monitors := rg.Group("/users/")
	monitors.GET("test_get/:uid/*action", testGetUser)
	monitors.POST("test_post_form/", testPostUser)
	monitors.POST("test_post/", testPostUserData)
}

type User struct {
	Name string `form:"name" json:"name"`
	Age  int    `form:"age" json:"age" binding:"required"`
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
	fmt.Println(string(userJson))
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

// http://localhost:5000/api/users/test_get/1/eat/?age=12
func testGetUser(c *gin.Context) {

	name := c.DefaultQuery("name", "pitou")
	age := c.Query("age")
	uid := c.Param("uid")
	action := c.Param("action")
	action = strings.Trim(action, "/")

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
