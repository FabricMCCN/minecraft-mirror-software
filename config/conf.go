package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"github.com/FabricMCCN/minecraft-mirror-software/constants"
	"github.com/FabricMCCN/minecraft-mirror-software/logger"
)

var once sync.Once
var ins *BaseConfig

func NewBaseConfig() *BaseConfig {
	once.Do(func() {
		ins = &BaseConfig{}
	})
	return ins
}

func (b *BaseConfig) LoadConfig() {
	defer func() {
		if e := recover(); e != nil {
			b.MinecraftAssets = constants.MinecraftAssets
			b.MinecraftLibraries = constants.MinecraftLibraries
			b.MinecraftVersionManifest = constants.MinecraftVersionManifest
			b.MinecraftLaunchMeta = constants.MinecraftLaunchMeta
			logger.Warn("An error occurred while reading the configuration file. Will use the default configuration file")
			logger.Warnf("%s", recover())
		}
	}()
	w, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	confPath := w + string(filepath.Separator) + "mirror.json"
	b1, err := ioutil.ReadFile(confPath)
	if err != nil {
		panic(err)
	}
	if json.Unmarshal(b1, b) != nil {
		panic(err)
	}
}
