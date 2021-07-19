package config

type BaseConfig struct {
	MinecraftLibraries       string `json:"minecraft_libraries,omitempty"`
	MinecraftAssets          string `json:"minecraft_assets,omitempty"`
	MinecraftVersionManifest string `json:"minecraft_version_manifest,omitempty"`
	MinecraftLaunchMeta      string `json:"minecraft_launch_meta,omitempty"`
}
