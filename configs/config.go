package configs

type ConfigDatabase struct {
	DriverName       string `mapstructure:"DRIVER_NAME"`
	ConnectionSource string `mapstructure:"CONNECTION_SOURCE"`
}

type ConfigRedis struct {
	Host   string `mapstructure:"HOST"`
	Port   int    `mapstructure:"PORT"`
	DB     int    `mapstructure:"db"`
	Pass   string `mapstructure:"pass"`
	Prefix string `mapstructure:"prefix"`
}

type ConfigCache struct {
	DriverName       string      `mapstructure:"DRIVER_NAME"`
	ConnectionSource string      `mapstructure:"CONNECTION_SOURCE"`
	Redis            ConfigRedis `mapstructure:"REDIS"`
}

type Config struct {
	AppName  string         `mapstructure:"APP_NAME"`
	WebPort  string         `mapstructure:"WEB_PORT"`
	Database ConfigDatabase `mapstructure:"DATABASE"`
	Cache    ConfigCache    `mapstructure:"CACHE"`
}
