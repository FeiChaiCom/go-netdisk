package user

import (
	"github.com/gin-gonic/gin"
	"go-netdisk/pkg/db/models"
	"go-netdisk/pkg/services/form"
	"go-netdisk/pkg/utils"
	"net/http"
)

// curl http://localhost:5000/api/user/page/?page=1&pageSize=20&orderCreateTime=DESC
func PageHandler(c *gin.Context) {
	var p form.UserParam
	if err := c.ShouldBindQuery(&p); err != nil {
		utils.Error(c, err)
		return
	}

	users, totalItems, totalPages := models.GetAllUsers(p.Page, p.PageSize, p.OrderCreateTime)

	utils.Ok(c, gin.H{
		"totalPage":  totalPages,
		"totalItems": totalItems,
		"data":       users,
	})
}

func Me(c *gin.Context) {
	username := c.GetString("username")
	me, err := models.GetUserByName(username)
	if err != nil {
		utils.Error(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": me,
	})
}
