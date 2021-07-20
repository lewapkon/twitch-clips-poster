package config

import (
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	TwitchTracker struct {
		Language    string `mapstructure:"language"`
		PagesToLoad int    `mapstructure:"pages_to_load"`
	} `mapstructure:"twitch_tracker"`
}

func ReadConfig() (Config, error) {
	config := Config{}
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("tcp")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	err := viper.ReadInConfig()
	if err != nil {
		return config, err
	}
	err = viper.Unmarshal(&config)
	return config, err
}
