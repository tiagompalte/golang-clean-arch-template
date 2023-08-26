package config

import (
	"os"

	"github.com/tiagompalte/golang-clean-arch-template/configs"
)

func ProviderSet() configs.Config {
	viper := NewViperConfig()

	path := os.Getenv("CONFIG_FILE")
	if path == "" {
		path = "./configs"
	}

	cfg, err := viper.Load(configs.ViperConfigurationName, configs.ViperTomlConfigurationType, path)
	if err != nil {
		panic(err)
	}

	return cfg
}
