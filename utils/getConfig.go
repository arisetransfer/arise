package utils

import (
	"os"

	"github.com/BurntSushi/toml"
)

type tomlConfig struct {
	Ip   string
	Port string
}

func GetIPAddrAndPort() (string, string) {
	var config tomlConfig
	var home string = os.Getenv("HOME")
	if _, err := toml.DecodeFile(home+"/.arise/config.toml", &config); err != nil {
		panic(err)
	}
	return config.Ip, config.Port
}
