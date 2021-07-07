package config

import (
	"fmt"
	"github.com/spf13/viper"
	"go-netdisk/utils"
	"log"
)

var ENV = &YamlConfig{}

const (
	MatterRootUUID     = "root"
	StaticDir          = "./static"
	StaticURL          = "/static"
	MediaURL           = "/media"
	TemplateDirPattern = "./templates/*"
	SimpleTime         = "2006-01-02 15:04:05"
)

func init() {
	// Load settings from config file or env
	LoadSettings("config")
	if err := viper.Unmarshal(ENV); err != nil {
		panic(err)
	}

	ENV.MatterRoot = ENV.MediaDir + "/matter-root"
	// viper.Debug()
	log.Println(utils.PrettyJson(ENV))
}

type LoginConfig struct {
	Ticket      string `mapstructure:"ticket" yaml:"ticket"`
	LoginURL    string `mapstructure:"login-url" yaml:"login-url"`
	UID         string `mapstructure:"uid" yaml:"uid"`
	UserInfoURL string `mapstructure:"user-info-url" yaml:"user-info-url"`
	SubPath     string `mapstructure:"sub-path" yaml:"sub-path"`
}

type YamlConfig struct {
	Port            int         `mapstructure:"port" yaml:"port"`
	Debug           bool        `mapstructure:"debug" yaml:"debug"`
	DefaultPassword string      `mapstructure:"default-password" yaml:"default-password"`
	LogFile         string      `mapstructure:"logfile" yaml:"logfile"`
	MediaDir        string      `mapstructure:"mediadir" yaml:"mediadir"`
	MatterRoot      string      `mapstructure:"matterroot" yaml:"matterroot"`
	Mysql           MysqlConfig `mapstructure:"mysql" yaml:"mysql"`
	JWT             JwtConfig   `mapstructure:"jwt" yaml:"jwt"`
	Login           LoginConfig `mapstructure:"login" yaml:"login"`
}

type JwtConfig struct {
	Issuer         string `mapstructure:"issuer" yaml:"issuer"`
	SecretKey      string `mapstructure:"secret-key"  yaml:"secret-key"`
	AuthCookieName string `mapstructure:"auth-cookie-name"  yaml:"auth-cookie-name"`
}

type MysqlConfig struct {
	Name     string `mapstructure:"name" yaml:"name"`
	Host     string `mapstructure:"host" yaml:"host"`
	Port     int    `mapstructure:"port"  yaml:"port"`
	Username string `mapstructure:"username" yaml:"username"`
	Password string `mapstructure:"password"  yaml:"password"`
}

func setDefaultSettings() {
	// viper.SetDefault("MatterRootUUID", "root")
}

func LoadSettings(fileName string) {
	log.Println("load system settings...")

	viper.AddConfigPath(".")
	viper.SetConfigName(fileName)
	viper.SetConfigType("yaml")

	// Auto get config from env
	viper.AutomaticEnv()

	// Read default config from yaml file
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Load config file error: %w \n", err))
	}

	// Write default settings
	setDefaultSettings()
}
