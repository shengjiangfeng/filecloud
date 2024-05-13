package service

import (
	"github.com/BurntSushi/toml"
)

type Storage struct {
}

type Config struct {
	WebAddr          string `toml:"WebAddr"`
	WebIndex         string `toml:"WebIndex"`
	StaticFS         bool   `json:"staticFS"`
	FilePath         string `toml:"FilePath"`
	FileDiskTotal    uint64 `toml:"FileDiskTotal"`
	SaveFileMultiple bool   `toml:"SaveFileMultiple"`
	Username         string `toml:"Username"`
	Password         string `toml:"Password"`
}

var config *Config

func LoadConfig(path string) *Config {
	conf := &Config{}
	_, err := toml.DecodeFile(path, conf)
	if err != nil {
		panic(err)
	}
	config = conf
	logger.Info(config)
	return config
}
