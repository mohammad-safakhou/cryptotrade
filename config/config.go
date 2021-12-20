package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	AppName       string         `mapstructure:"appname"`
	Database      DatabaseConfig `mapstructure:"database"`
	Env           EnvConfig      `mapstructure:"env"`
	Server        ServerConfig   `mapstructure:"server"`
	MessageBroker MessageBroker  `mapstructure:"message_broker"`
}

type DatabaseConfig struct {
	PostgresConfig PostgresConfig `mapstructure:"psql"`
}

type PostgresConfig struct {
	Type     string `mapstructure:"type"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
	SslMode  string `mapstructure:"sslmode"`
}

type EnvConfig struct {
}

type ServerConfig struct {
	Host  string `mapstructure:"host"`
	Port  string `mapstructure:"port"`
	Debug bool   `mapstructure:"debug"`
}

type MessageBroker struct {
	Kafka Kafka `mapstructure:"kafka"`
}

type Kafka struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

func ViperConfig() (*viper.Viper, Config, error) {
	v := viper.New()
	v.AutomaticEnv()

	v.SetConfigName("config")
	v.AddConfigPath("./config")
	v.AddConfigPath("/config")
	v.AddConfigPath("/app/config")
	err := v.ReadInConfig()
	if err != nil {
		return nil, Config{}, err
	}
	var config Config
	err = v.Unmarshal(&config)
	if err != nil {
		return nil, Config{}, err
	}

	return v, config, nil
}
