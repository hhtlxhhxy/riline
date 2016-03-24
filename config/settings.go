package config

import (
	"github.com/go-ini/ini"
	"fmt"
)

var Configs *RilineSettings

type RilineSettings struct {
	ProfilingSetting Profiling
	MySQLDBSetting   MySQLDB
}

type Profiling struct {
	Enable bool
	Host   string
	Port   string
}

type MySQLDB struct {
	Host       string
	Port       string
	Database   string
	User       string
	Password   string
	ScriptFile string
}

func InitConfig(config string) (err error) {
	var cfg *ini.File
	cfg, err = ini.Load(config)
	if err != nil {
		fmt.Println("Read config file error: " + config)
		return err
	}
	cfg.NameMapper = ini.TitleUnderscore

	Configs = new(RilineSettings)
	err = cfg.Section("profiling").MapTo(&Configs.ProfilingSetting)
	if err != nil {
		return err
	}
	err = cfg.Section("mysql").MapTo(&Configs.MySQLDBSetting)
	if err != nil {
		return err
	}
	return nil
}
