package common

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type Config struct {
	Database DbConfig
}

type DbConfig struct {
	Name     string `toml:"name"`
	Port     int    `toml:"port"`
	Username string `toml:"username"`
	Password string `toml:"password"`
}

var config *Config

func LoadConfig(path string) (*Config, error) {
	_, err := toml.DecodeFile(path, &config)
	return config, err
}

func (config *Config) ConnectString() string {
	return fmt.Sprintf("%s:%s@/%s?charset=utf8",
		config.Database.Username,
		config.Database.Password,
		config.Database.Name)
}
