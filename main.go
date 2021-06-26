package main

import (
	"fmt"
	"github.com/gaomugong/go-netdisk/apps"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	port := 5000
	fmt.Println("go-netdisk begin server at port: ", port)

	// Init url router for apis
	router := apps.InitApiRouter()

	// Load index html
	router.LoadHTMLGlob("templates/*")
	router.GET("", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "feichai",
		})
	})

	// Serve static files
	router.Static("/static", "./statics")
	//router.StaticFS("/more_static", http.Dir("my_file_system"))
	router.StaticFile("/favicon.ico", "./statics/favicon.ico")

	_ = router.Run(fmt.Sprintf("127.0.0.1:%d", port))
}
