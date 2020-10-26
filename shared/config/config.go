package config

import (
	"os"

	"github.com/spf13/viper"
)

type (
	// AppConfigInterface is represent method in appConfig structs
	AppConfigInterface interface {
		GetAppPort() int
		GetDbHost() string
		GetDbPort() int
		GetDbUser() string
		GetDbPassword() string
		GetDbName() string
	}

	// appConfig is a struct to mapping app configuration object
	appConfig struct {
		AppPort    int    `mapstructure:"APP_PORT"`
		DbHost     string `mapstructure:"DB_HOST"`
		DbPort     int    `mapstructure:"DB_PORT"`
		DbUser     string `mapstructure:"DB_USER"`
		DbPassword string `mapstructure:"DB_PASSWORD"`
		DbName     string `mapstructure:"DB_NAME"`
	}
)

func (c *appConfig) GetAppPort() int {
	return c.AppPort
}

func (c *appConfig) GetDbHost() string {
	return c.DbHost
}

func (c *appConfig) GetDbPort() int {
	return c.DbPort
}

func (c *appConfig) GetDbUser() string {
	return c.DbUser
}

func (c *appConfig) GetDbPassword() string {
	return c.DbPassword
}

func (c *appConfig) GetDbName() string {
	return c.DbName
}

var (
	runtimeConfig appConfig
)

// NewConfig return configurations implementation
func NewConfig() (AppConfigInterface, error) {
	v := viper.New()

	appEnv, exists := os.LookupEnv("APP_ENV")
	if exists {
		if appEnv == "staging" {
			v.SetConfigFile("app.config.staging")
		} else if appEnv == "production" {
			v.SetConfigName("app.config.prod")
		}
	} else {
		v.SetConfigName("app.config.dev")
	}

	v.AddConfigPath(".")

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	v.Unmarshal(&runtimeConfig)

	return &runtimeConfig, nil
}
