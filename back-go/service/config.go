package service

import (
	"github.com/BurntSushi/toml"
)

type Storage struct {
	Provider     string `toml:"Provider"` // aws google aliyun tencent
	Bucket       string `toml:"Bucket"`
	AccessKey    string `toml:"AccessKey"`
	AccessSecret string `toml:"AccessSecret"`
	Endpoint     string `toml:"endpoint"`
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
	StorageConfig    string `toml:"StorageConfig"`
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
