package util

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Port                 string        `mapstructure:"PORT"`
	TokenSymmetricKey    string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	PostgresURL          string        `mapstructure:"POSTGRES_URL"`
	MigrationURL         string        `mapstructure:"MIGRATION_URL"`
	MaxConns             int           `mapstructure:"MAX_CONNS"`
	LogLevel             string        `mapstructure:"LOG_LEVEL"`
}

func LoadConfig(path, fileName, configType string) (*Config, error) {
	v := viper.New()

	v.AddConfigPath(path)
	v.SetConfigName(fileName)
	v.SetConfigType(configType)

	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	conf := &Config{}
	if err := v.Unmarshal(&conf); err != nil {
		return nil, err
	}

	return conf, nil
}
