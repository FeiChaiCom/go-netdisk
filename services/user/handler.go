package user

import (
	"github.com/gin-gonic/gin"
	"go-netdisk/models/db"
	"go-netdisk/models/form"
	R "go-netdisk/render"
)

// curl http://localhost:5000/api/user/page/?page=1&pageSize=20&orderCreateTime=DESC
func PageHandler(c *gin.Context) {
	var p form.UserParam
	if err := c.ShouldBindQuery(&p); err != nil {
		R.Error(c, err)
		return
	}

	users, totalItems, totalPages := db.GetAllUsers(p.Page, p.PageSize, p.OrderCreateTime)

	R.Ok(c, gin.H{
		"totalPage":  totalPages,
		"totalItems": totalItems,
		"data":       users,
	})
}
