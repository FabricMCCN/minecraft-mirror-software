package providers

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"github.com/FabricMCCN/minecraft-mirror-software/constants"
	"github.com/FabricMCCN/minecraft-mirror-software/jsonstrs"
	"github.com/FabricMCCN/minecraft-mirror-software/utils"
)

var mOnce sync.Once
var mIns *ManifestsProvider

type ManifestsProvider struct {
	DownloadInfo []utils.DownloadInfo
}

func NewManifestsProvider() *ManifestsProvider {
	mOnce.Do(func() { mIns = &ManifestsProvider{} })
	return mIns
}

func (m *ManifestsProvider) ParseManifests(b []byte) {
	t := &jsonstrs.Manifest{}
	json.Unmarshal(b, t)
	for _, v := range t.Versions {
		m.DownloadInfo = append(m.DownloadInfo, utils.DownloadInfo{
			DownloadUrl:     strings.ReplaceAll(v.URL, constants.MinecraftLaunchMeta, utils.MirrorUtil{}.LaunchMetaURL()),
			DownloadStorage: storagePath(v.URL),
			Hash:            v.Sha1,
			Size:            0,
		})
	}
}

func storagePath(url string) string {
	wd, _ := os.Getwd()
	if runtime.GOOS == "windows" {
		return filepath.FromSlash(wd + string(filepath.Separator) + strings.ReplaceAll(url, constants.MinecraftLaunchMeta, ""))
	}
	return wd + string(filepath.Separator) + strings.ReplaceAll(url, constants.MinecraftLaunchMeta, "")
}
