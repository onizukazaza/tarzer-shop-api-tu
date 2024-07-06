package config

import (
	"strings"
	"sync"
	"time"

	"github.com/go-playground/validator"
	"github.com/spf13/viper"
)

type (
	Config struct {
		Server   *Server   `mapstructure:"server" validate:"required"`
		OAuth2   *OAuth2   `mapstructure:"oauth2" validate:"required"`
		State    *State    `mapstructure:"state" validate:"required"`
		Database *Database `mapstructure:"database" validate:"required"`
	}
	Server struct {
		Port         int           `mapstructure:"port" validate:"required"` //no put is error
		AllowOrigins []string      `mapstructure:"allowOrigins" validate:"required"`
		BodyLimit    string        `mapstructure:"bodyLimit" validate:"required"`
		TimeOut      time.Duration `mapstructure:"timeout" validate:"required"`
	}
	OAuth2 struct {
		PlayerRedirectUrl string   `mapstructure:"playerRedirectUrl" validate:"required"`
		AdminRedirectUrl  string   `mapstructure:"adminRedirectUrl" validate:"required"`
		ClientID          string   `mapstructure:"clientID" validate:"required"`
		ClientSecret      string   `mapstructure:"clientSecret" validate:"required"`
		Endpoints         endpoint `mapstructure:"endpoints" validate:"required"`
		Scopes            []string `mapstructure:"scopes" validate:"required"`
		UrlInfoUrl        string   `mapstructure:"uelInfoUrl" validate:"required"`
		RevokeUrl         string   `mapstructure:"revokeUrl" validate:"required"`
	}

	endpoint struct {
		AuthUrl       string `mapstructure:"authUrl" validate:"required"`
		TokenUrl      string `mapstructure:"tokenUrl" validate:"required"`
		DeviceAuthUrl string `mapstructure:"deviceAuthUrl" validate:"required"`
	}
	State struct {
		Secret    string        `mapstructure:"secret" validate:"required"`
		ExpiresAT time.Duration `mapstructure:"expiresAT" validate:"required"`
		Issuer    string        `mapstructure:"issuer" validate:"required"`
	}
	Database struct {
		Host     string `mapstructure:"host" validate:"required"`
		Port     int    `mapstructure:"port" validate:"required"`
		User     string `mapstructure:"user" validate:"required"`
		Password string `mapstructure:"password" validate:"required"`
		Dbname   string `mapstructure:"dbname" validate:"required"`
		Sslmode  string `mapstructure:"sslmode" validate:"required"`
		Schema   string `mapstructure:"schema" validate:"required"`
	}
)

var (
	once           sync.Once
	configInstance *Config
)

func ConfigGetting() *Config {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./config")
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) //server.port -> server_port

		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}
		if err := viper.Unmarshal(&configInstance); err != nil {
			panic(err)
		}

		validating := validator.New() 

		if err := validating.Struct(configInstance); err != nil { // & viper base on
			panic(err)
		}
	})
	return configInstance
}
