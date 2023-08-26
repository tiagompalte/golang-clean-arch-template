package config

import (
	"github.com/spf13/viper"
	"github.com/tiagompalte/golang-clean-arch-template/configs"
)

type ViperConfig struct{}

func NewViperConfig() Config {
	return ViperConfig{}
}

func (v ViperConfig) Load(filename string, configType string, path string) (configs.Config, error) {
	viper.SetConfigName(filename)
	viper.SetConfigType(configType)
	viper.AddConfigPath(path)
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return configs.Config{}, err
	}

	var config configs.Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return configs.Config{}, err
	}

	return config, nil
}
