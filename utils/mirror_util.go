package utils

import (
	"github.com/FabricMCCN/minecraft-mirror-software/config"
	"github.com/FabricMCCN/minecraft-mirror-software/constants"
)

type MirrorUtil struct {
}

func (m MirrorUtil) LibrariesURL() string {
	if config.NewBaseConfig().MinecraftLibraries == "" {
		return constants.MinecraftLibraries
	}
	return config.NewBaseConfig().MinecraftLibraries
}

func (m MirrorUtil) AssetsURL() string {
	if config.NewBaseConfig().MinecraftAssets == "" {
		return constants.MinecraftAssets
	}
	return config.NewBaseConfig().MinecraftAssets
}

func (m MirrorUtil) VersionManifestURL() string {
	if config.NewBaseConfig().MinecraftVersionManifest == "" {
		return constants.MinecraftVersionManifest
	}
	return config.NewBaseConfig().MinecraftVersionManifest
}

func (m MirrorUtil) LaunchMetaURL() string {
	if config.NewBaseConfig().MinecraftLaunchMeta == "" {
		return constants.MinecraftLaunchMeta
	}
	return config.NewBaseConfig().MinecraftLaunchMeta
}
