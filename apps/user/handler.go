package user

import (
	"github.com/gaomugong/go-netdisk/models/db"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type userParam struct {
	Page            int    `form:"page"`
	PageSize        int    `form:"pageSize"`
	OrderCreateTime string `form:"orderCreateTime"`
}

// curl http://localhost:5000/api/account/users/?page=1&pageSize=20&orderCreateTime=DESC
func PageHandler(c *gin.Context) {
	var p userParam
	if err := c.ShouldBindQuery(&p); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":  false,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	users, totalItems, totalPages := db.GetAllUsers(p.Page, p.PageSize, p.OrderCreateTime)
	log.Printf("%#v %d %d\n", p, totalItems, totalPages)

	c.JSON(http.StatusOK, gin.H{
		"result": true,
		"data": gin.H{
			"totalPage":  totalPages,
			"totalItems": totalItems,
			"data":       users,
		},
		"message": "ok",
	})
}
