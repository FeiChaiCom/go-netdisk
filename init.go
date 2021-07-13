package main

import (
	"github.com/gin-gonic/gin"
	cfg "go-netdisk/config"
	"go-netdisk/models/db"
	"go-netdisk/utils"
	"io"
	"log"
	"os"
)

// Init gin log to file and stdout
func initServer() {
	log.Println("init gin log to gin.log and stdout...")
	f, _ := os.Create(cfg.ENV.LogFile)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	log.Println("init file upload dir...")
	if _, err := os.Stat(cfg.ENV.MediaDir); os.IsNotExist(err) {
		if err = os.Mkdir(cfg.ENV.MediaDir, 0755); err != nil {
			panic(err)
		}
	}

	if _, err := os.Stat(cfg.ENV.MatterRoot); os.IsNotExist(err) {
		if err = os.Mkdir(cfg.ENV.MatterRoot, 0755); err != nil {
			panic(err)
		}
	}

	if !cfg.ENV.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	// Init mysql connection
	if err := cfg.InitDB(); err != nil {
		panic(err)
	}

	// Initialize database
	initDatabase()
}

func initDatabase() {
	if cfg.ENV.NeedMigrate {
		_ = cfg.DB.AutoMigrate(&db.Project{}, &db.User{}, &db.Permission{}, &db.Matter{}, db.Preference{})
	}

	log.Printf("Create superuser: %s", cfg.ENV.SuperUser)
	if _, err := db.GetOrCreateUser(cfg.ENV.SuperUser, true); err != nil {
		panic(err)
	}

	perm := &db.Permission{}
	cfg.DB.Where(db.Permission{UserName: cfg.ENV.SuperUser}).Attrs(db.Permission{
		Role: db.ADMINISTRATOR,
	}).FirstOrCreate(&perm)
	log.Printf("GetOrCreate permission: %s\n", utils.PrettyJson(perm))

	prefer := &db.Preference{}
	cfg.DB.Where(db.Preference{Name: "netdisk"}).Attrs(db.Preference{
		AllowRegister: true,
	}).FirstOrCreate(&prefer)
	log.Printf("GetOrCreate preference: %s\n", utils.PrettyJson(prefer))

	if cfg.DB.First(&db.Project{}).RowsAffected == 0 {
		log.Printf("Create default project")
		project := db.Project{
			Name:        "DEMO",
			Description: "DEMO",
		}
		if err := cfg.DB.Create(&project).Error; err != nil {
			panic(err)
		}

		guest, _ := db.GetOrCreateUser("user0", false)
		cfg.DB.Create(&db.Permission{
			UserName:    guest.Username,
			ProjectUUID: project.UUID,
			Role:        "USER",
		})
	}

}
