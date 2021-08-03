package demo

import (
	"github.com/gin-gonic/gin"
)

// Add demo apis to api group
func RegisterTestGroup(rg *gin.RouterGroup) {
	users := rg.Group("/tests/")

	// if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
	// 	fmt.Println("bind register of demo")
	// 	_ = v.RegisterValidation("birthDayValidator", birthDayValidator)
	// }

	users.GET("test_get/:uid/*action", testGetUser)
	users.GET("test_redirect/", testRedirect)
	users.POST("test_post_form/", testPostUser)
	users.POST("test_post/", testPostUserData)
	users.POST("test_upload/", testUploadFile)
	users.GET("test_get_file/", testGetFile)

}
