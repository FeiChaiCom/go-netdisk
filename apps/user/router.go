package user

import (
	"github.com/gin-gonic/gin"
)

// Add user apis to api group
func RegisterUserGroup(rg *gin.RouterGroup) {
	users := rg.Group("/users/")

	//if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
	//	fmt.Println("bind register of user")
	//	_ = v.RegisterValidation("birthDayValidator", birthDayValidator)
	//}
	users.GET("test_get/:uid/*action", testGetUser)
	users.GET("test_redirect/", testRedirect)
	users.POST("test_post_form/", testPostUser)
	users.POST("test_post/", testPostUserData)

}
