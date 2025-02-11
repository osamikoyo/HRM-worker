package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct{
	Port string `mapstructure:"PORT"`
	Host string `mapstructure:"HOST"`
    DatabaseURL string `mapstructure:"TURSO_URL"`
}

func LoadConfig() (c Config, err error) {
    viper.AddConfigPath("./internal/config/envs")
    viper.SetConfigName("dev")
    viper.SetConfigType("env")

    viper.AutomaticEnv()

    err = viper.ReadInConfig()

    if err != nil {
        return
    }

    err = viper.Unmarshal(&c)

    c.DatabaseURL = os.Getenv("TURSO_URL")

    return
}