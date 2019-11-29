package main

import (
	"github.com/BurntSushi/toml"
	"log"
)

var (
	DbCfg DatabaseConfig
)

func init() {
	fileCfg := loadFileConfig()
	DbCfg = fileCfg.Database
}

func loadFileConfig() FileConfig {
	var fileCfg FileConfig
	if _, err := toml.DecodeFile("./config.toml", &fileCfg); err != nil {
		log.Fatal(err)
	}
	return fileCfg
}

type DatabaseConfig struct {
	Host          string
	Port          int
	DbName        string
	Usr           string
	Psw           string
	MaxConnection int
	Enabled       bool
}

type FileConfig struct {
	Database DatabaseConfig `toml:"Database"`
}
