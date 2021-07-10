package config

import (
	"fmt"
	"github.com/spf13/viper"
	"go-netdisk/utils"
	"log"
	"os"
)

var (
	RunMode = "dev"
	ENV     = &YamlConfig{}
)

const (
	MatterRootUUID     = "root"
	StaticDir          = "./static"
	StaticURL          = "/static"
	MediaURL           = "/media"
	TemplateDirPattern = "./templates/*"
	SimpleTime         = "2006-01-02 15:04:05"
)

func init() {
	// dev/stag/prod.yaml
	if RunMode = os.Getenv("BKPAAS_ENVIRONMENT"); RunMode == "" {
		RunMode = "dev"
	}

	// Load settings from config file or env
	LoadSettings(RunMode)
	if err := viper.Unmarshal(ENV); err != nil {
		panic(err)
	}

	// viper.Debug()
	log.Println(utils.PrettyJson(ENV))
}

type YamlConfig struct {
	Debug   bool   `mapstructure:"debug"`
	RunMode string `mapstructure:"runmode"`
	Port    int    `mapstructure:"port"`
	LogFile string `mapstructure:"logfile"`

	Mysql MysqlConfig `mapstructure:"mysql"`
	JWT   JwtConfig   `mapstructure:"jwt"`
	Login LoginConfig `mapstructure:"login"`
	Paas  PaasConfig  `mapstructure:"paas"`

	// service logic related config
	DefaultPassword string `mapstructure:"default-password"`
	MediaDir        string `mapstructure:"media-dir"`
	MatterRoot      string `mapstructure:"upload-dir"`
}

type JwtConfig struct {
	Issuer         string `mapstructure:"issuer"`
	SecretKey      string `mapstructure:"secret-key"`
	AuthCookieName string `mapstructure:"auth-cookie-name"`
}

type LoginConfig struct {
	UID         string `mapstructure:"uid"`
	Ticket      string `mapstructure:"ticket"`
	LoginURL    string `mapstructure:"login-url"`
	UserInfoURL string `mapstructure:"user-info-url"`
	SubPath     string `mapstructure:"sub-path"`
}

type PaasConfig struct {
	URL           string `mapstructure:"url"`
	AppName       string `mapstructure:"app-name"`
	AppCode       string `mapstructure:"app-code"`
	AppSecret     string `mapstructure:"app-secret"`
	AppModuleName string `mapstructure:"app-module-name"`
	AppLogPath    string `mapstructure:"app-log-path"`
}

type MysqlConfig struct {
	Name     string `mapstructure:"name"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

func bindEnvSettings() {
	_ = viper.BindEnv("port", "PORT")
	_ = viper.BindEnv("runmode", "BKPAAS_ENVIRONMENT")

	_ = viper.BindEnv("mysql.name", "GCS_MYSQL_NAME")
	_ = viper.BindEnv("mysql.host", "GCS_MYSQL_HOST")
	_ = viper.BindEnv("mysql.port", "GCS_MYSQL_PORT")
	_ = viper.BindEnv("mysql.username", "GCS_MYSQL_USER")
	_ = viper.BindEnv("mysql.password", "GCS_MYSQL_PASSWORD")

	_ = viper.BindEnv("paas.url", "BKPAAS_URL")
	_ = viper.BindEnv("paas.app-name", "BKPAAS_ENGINE_APP_NAME")
	_ = viper.BindEnv("paas.app-code", "BKPAAS_APP_CODE")
	_ = viper.BindEnv("paas.app-secret", "BKPAAS_APP_SECRET")
	_ = viper.BindEnv("paas.app-module-name", "BKPAAS_APP_MODULE_NAME")
	_ = viper.BindEnv("paas.app-log-path", "BKPAAS_APP_LOG_PATH")

	_ = viper.BindEnv("login.ticket", "BKAPP_TICKET")
	_ = viper.BindEnv("login.login-url", "BKAPP_LOGIN_URL")
	_ = viper.BindEnv("login.user-info-url", "BKAPP_USER_INFO_URL")
	_ = viper.BindEnv("login.uid", "BKAPP_UID")

}

func setDefaultSettings() {
	viper.SetDefault("runmode", "dev")
}

func LoadSettings(fileName string) {
	log.Printf("load settings for <%s> ...\n", RunMode)

	viper.AddConfigPath(".envs")
	viper.SetConfigName(fileName)
	viper.SetConfigType("yaml")

	// Auto get config from env
	viper.AutomaticEnv()

	// Bind env config
	bindEnvSettings()

	// Write default settings
	setDefaultSettings()

	// Read default config from yaml file
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Load config file error: %w \n", err))
	}

}
