package providers

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/FabricMCCN/minecraft-mirror-software/config"
	"github.com/FabricMCCN/minecraft-mirror-software/jsonstrs"
	"github.com/FabricMCCN/minecraft-mirror-software/utils"
)

var vOnce sync.Once
var vIns *LibrariesProvider

type LibrariesProvider struct {
	DownloadInfo []utils.DownloadInfo
}

func NewLibrariesProvider() *LibrariesProvider {
	vOnce.Do(func() { vIns = &LibrariesProvider{} })
	return vIns
}

func (ve *LibrariesProvider) Parse(b []byte) error {
	version := &jsonstrs.VersionDefine{}

	json.Unmarshal(b, version)
	for _, v := range version.Libraries {
		ve.DownloadInfo = append(ve.DownloadInfo, utils.DownloadInfo{
			DownloadUrl:     config.NewBaseConfig().MinecraftLibraries + *v.Downloads.Artifact.Path,
			Hash:            v.Downloads.Artifact.Sha1,
			DownloadStorage: storageLibrariesPath(*v.Downloads.Artifact.Path),
		})
		if v.Downloads.Classifiers != nil {
			if v.Downloads.Classifiers.NativesLinux != nil {
				ve.DownloadInfo = append(ve.DownloadInfo, utils.DownloadInfo{
					DownloadUrl:     config.NewBaseConfig().MinecraftLibraries + *v.Downloads.Classifiers.NativesLinux.Path,
					Hash:            v.Downloads.Classifiers.NativesLinux.Sha1,
					DownloadStorage: storageLibrariesPath(*v.Downloads.Classifiers.NativesLinux.Path),
				})
			}
			if v.Downloads.Classifiers.NativesMacos != nil {
				ve.DownloadInfo = append(ve.DownloadInfo, utils.DownloadInfo{
					DownloadUrl:     config.NewBaseConfig().MinecraftLibraries + *v.Downloads.Classifiers.NativesMacos.Path,
					Hash:            v.Downloads.Classifiers.NativesMacos.Sha1,
					DownloadStorage: storageLibrariesPath(*v.Downloads.Classifiers.NativesMacos.Path),
				})
			}
			if v.Downloads.Classifiers.NativesWindows != nil {
				ve.DownloadInfo = append(ve.DownloadInfo, utils.DownloadInfo{
					DownloadUrl:     config.NewBaseConfig().MinecraftLibraries + *v.Downloads.Classifiers.NativesWindows.Path,
					Hash:            v.Downloads.Classifiers.NativesWindows.Sha1,
					DownloadStorage: storageLibrariesPath(*v.Downloads.Classifiers.NativesWindows.Path),
				})
			}
			if v.Downloads.Classifiers.NativesOsx != nil {
				ve.DownloadInfo = append(ve.DownloadInfo, utils.DownloadInfo{
					DownloadUrl:     config.NewBaseConfig().MinecraftLibraries + *v.Downloads.Classifiers.NativesOsx.Path,
					Hash:            v.Downloads.Classifiers.NativesOsx.Sha1,
					DownloadStorage: storageLibrariesPath(*v.Downloads.Classifiers.NativesOsx.Path),
				})
			}
		}

	}

	return nil
}

func (v *LibrariesProvider) GetDownloadList() []utils.DownloadInfo {
	return v.DownloadInfo
}

func storageLibrariesPath(path string) string {
	wd, _ := os.Getwd()
	p := wd + string(filepath.Separator) + "libraries" + string(filepath.Separator) + path
	if runtime.GOOS == "windows" {
		return filepath.FromSlash(p)
	}
	return p
}
