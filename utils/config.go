package utils

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Port        string `mapstructure:"PORT"`
	Version     string `mapstructure:"VERSION"`
	DatabaseURL string `mapstructure:"DATABASE_URL"`
}

func LoadConfig() (*Config, error) {
	viper.SetDefault("PORT", "8080")
	viper.SetDefault("VERSION", "1.0.0")

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	viper.AutomaticEnv()
	viper.BindEnv("DATABASE_URL")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	if cfg.DatabaseURL == "" {
		return nil, fmt.Errorf("DATABASE_URL is required but not set")
	}

	return &cfg, nil
}
