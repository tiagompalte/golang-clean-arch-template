package configs

type ConfigDatabase struct {
	DriverName       string `mapstructure:"DRIVER_NAME"`
	ConnectionSource string `mapstructure:"CONNECTION_SOURCE"`
}

type ConfigCache struct {
	DriverName       string `mapstructure:"DRIVER_NAME"`
	ConnectionSource string `mapstructure:"CONNECTION_SOURCE"`
}

type Config struct {
	AppName   string         `mapstructure:"APP_NAME"`
	WebServer string         `mapstructure:"WEB_SERVER"`
	WebPort   string         `mapstructure:"WEB_PORT"`
	Database  ConfigDatabase `mapstructure:"DATABASE"`
	Cache     ConfigCache    `mapstructure:"CACHE"`
}
