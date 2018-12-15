package config

import (
	"gopkg.in/ini.v1"
)

type ConfigList struct {
	Port int
}

var Config ConfigList

func init() {
	cfg, _ := ini.Load("config.ini")
	Config = ConfigList{
		Port: cfg.Section("web").Key("port").MustInt(),
	}
}