package config

import "os"

// Config Global Config variable
var Config Configuration

type Configurer interface {
	Get() Configuration
}

type DefaultConfig struct{}

func (c *DefaultConfig) Get() Configuration {
	env := os.Getenv("ENV")

	if len(env) == 0 {
		env = "local"
	}

	switch env {
	case "local":
		return getLocalConfig()
		// case "sandbox":
		// 	return getDevConfig()
		// case "test":
		// 	fallthrough
		// case "production":
		// 	return getProdConfig()
	}

	return Configuration{}
}

func init() {
	InitConfig()
}

func InitConfig() {
	config := GetConfig(&DefaultConfig{})
	Config = config
}

func GetConfig(c Configurer) Configuration {
	return c.Get()
}
