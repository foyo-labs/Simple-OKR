package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var (
	C = new(Config)
)

type HttpConf struct {
	Port       string `mapstructure:"port"`
	XRequestID string
	Timeout    int
}

type Postgres struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"name"`
	SSLMode  string `mapstructure:"ssl"`
}

// Log 日志配置参数
type Log struct {
	Level  int
	Format string
	Output string
	File   string
}

func (a Postgres) DSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		a.Host, a.Port, a.User, a.DBName, a.Password, a.SSLMode)
}

type Config struct {
	RunMode  string   `mapstructure:"profile"`
	Http     HttpConf `mapstructure:"http"`
	Database Postgres `mapstructure:"postgres"`
	Log      Log      `mapstructure:"log"`
}

func LoadConfig() (err error) {
	viper.SetConfigFile("./config.yml")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(C)
	return
}
