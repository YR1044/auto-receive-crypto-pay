package internal

import (
	"github.com/BurntSushi/toml"
	"github.com/HEUDavid/go-fsm/pkg/util"
	"log"
	"path/filepath"
)

type Config struct {
	Global Global `toml:"global"`
}

type Global struct {
	LogPath      string   `toml:"logPath"`
	Addr         string   `toml:"addr"`
	AdminAddress []string `toml:"adminAddress"`
}

func GetConfig() Config {
	var config Config
	if _, err := toml.DecodeFile(filepath.Join(util.FindProjectRoot(), "conf", "conf.toml"), &config); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	return config
}