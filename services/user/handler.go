package user

import (
	"github.com/gaomugong/go-netdisk/models/db"
	"github.com/gaomugong/go-netdisk/models/form"
	R "github.com/gaomugong/go-netdisk/render"
	"github.com/gin-gonic/gin"
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
