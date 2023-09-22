package util

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Token    TokenConfig
	Database DatabaseConfig
	App      AppConfig
}

type ServerConfig struct {
	Port string `mapstructure:"PORT"`
}

type TokenConfig struct {
	TokenSymmetricKey    string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
}

type DatabaseConfig struct {
	PostgresURL  string `mapstructure:"POSTGRES_URL"`
	MigrationURL string `mapstructure:"MIGRATION_URL"`
	MaxConns     int    `mapstructure:"MAX_CONNS"`
}

type AppConfig struct {
	LogLevel    string `mapstructure:"LOG_LEVEL"`
	Environment string `mapstructure:"ENVIRONMENT"`
}

func LoadConfig(path, fileName, configType string) (*Config, error) {
	v := viper.New()

	v.SetConfigName(fileName)
	v.AddConfigPath(path)
	viper.SetConfigType(configType)

	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	conf := &Config{}
	if err := v.Unmarshal(conf); err != nil {
		return nil, err
	}

	return conf, nil
}
