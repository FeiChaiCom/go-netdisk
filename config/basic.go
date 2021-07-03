package config

const (
	Port               = 5000
	JwtIssuer          = "FEICHAI.COM"
	JwtSecretKey       = "feichaicom"
	AuthCookieName     = "auth_token"
	LogFile            = "gin.log"
	MediaDir           = "./media"
	MatterRootUUID     = "root"
	MatterRoot         = MediaDir + "/matter-root"
	StaticDir          = "./statics"
	StaticURL          = "/static"
	MediaURL           = "/media"
	TemplateDirPattern = "./templates/*"
)
