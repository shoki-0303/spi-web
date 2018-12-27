package config

import (
	"gopkg.in/ini.v1"
)

type ConfigList struct {
	Port       int
	DbName     string
	DbDriver   string
	LogFile    string
	DigitalKey string
}

var Config ConfigList

func init() {
	cfg, _ := ini.Load("config.ini")
	Config = ConfigList{
		Port:       cfg.Section("web").Key("port").MustInt(),
		DigitalKey: cfg.Section("web").Key("digitalKey").String(),
		DbName:     cfg.Section("db").Key("name").String(),
		DbDriver:   cfg.Section("db").Key("driver").String(),
		LogFile:    cfg.Section("spi").Key("logfile").String(),
	}
}
