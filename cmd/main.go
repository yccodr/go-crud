package main

import (
	"go-cleanarch/internal/router"
	"log/slog"

	"github.com/spf13/viper"
)

type Config struct {
	DB_DSN string `mapstructure:"DB_DSN"`
}

func readConfig(config *Config, path string) {
	viper.SetConfigType("yaml")
	viper.SetDefault("DB_DSN", "")
	viper.AutomaticEnv()

	if path != "" {
		viper.SetConfigFile(path)
	}

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			slog.Info("[Config] Config file not found. Reading from environment variables.")
		} else {
			slog.Error("[Config] Error reading config file", "err", err)
			panic(err)
		}
	}

	viper.Unmarshal(config)
}

func main() {
	readConfig(&Config{}, "")

	router.NewRouter().Run(":8080")
}
