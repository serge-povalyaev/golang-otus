package calendar

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	RepositoryType string
	Server         Server
	GRPC           GRPC
	DB             DBConfig
	Logger         LoggerConfig
}

type Server struct {
	Host string
	Port string
}

type GRPC struct {
	Host string
	Port string
}

type DBConfig struct {
	Host string
	Port string
	User string
	Pass string
	Name string
}

type LoggerConfig struct {
	Level    string
	FilePath string
}

func ReadConfig(configPath string) Config {
	viper.SetConfigFile(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Fatal(fmt.Sprintf("Не удалось прочитать файл конфигурации: %s", err))
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		logrus.Fatal(fmt.Sprintf("Не удалось получить конфигурацию: %s", err))
	}

	return config
}
