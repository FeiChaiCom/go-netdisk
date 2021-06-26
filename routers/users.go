package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// Add user apis to api group
func registerUserGroup(rg *gin.RouterGroup) {
	monitors := rg.Group("/users/")
	monitors.GET("test_get/:uid/*action", testGetUser)
	monitors.POST("test_post/", testPostUser)
}

//curl http://localhost:5000/api/users/test_post/ -X POST -F name=miya -F age=12
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
