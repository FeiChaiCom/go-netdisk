package main

import (
	"fmt"
	"github.com/gaomugong/go-netdisk/apps"
	cfg "github.com/gaomugong/go-netdisk/config"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// Init url router for apis
	router := apps.InitAPIRouter()

	// Load index html
	router.LoadHTMLGlob(cfg.TemplateDirPattern)
	router.GET("", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "feichai",
		})
	})

	// Serve static files
	router.Static(cfg.StaticURL, cfg.StaticDir)
	router.StaticFile("/favicon.ico", fmt.Sprintf("%s/favicon.ico", cfg.StaticDir))

	// Serve media files
	router.StaticFS(cfg.MediaURL, http.Dir(cfg.MediaDir))

	_ = router.Run(fmt.Sprintf(":%d", cfg.Port))
}

// Init gin log to file and stdout
func init() {
	log.Println("init gin log to gin.log and stdout...")
	f, _ := os.Create(cfg.LogFile)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	log.Println("init file upload dir...")
	if _, err := os.Stat(cfg.MediaDir); os.IsNotExist(err) {
		if err = os.Mkdir(cfg.MediaDir, 0755); err != nil {
			panic(err)
		}
	}
}
