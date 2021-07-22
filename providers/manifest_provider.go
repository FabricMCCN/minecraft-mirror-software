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
	"github.com/FabricMCCN/minecraft-mirror-software/logger"
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

func (m *ManifestsProvider) GetDownloadList() []utils.DownloadInfo { return m.DownloadInfo }

func (m *ManifestsProvider) Parse(b []byte) error {
	t := &jsonstrs.Manifest{}
	if err := json.Unmarshal(b, t); err != nil {
		logger.Error(err)
		return err
	}
	for _, v := range t.Versions {
		m.DownloadInfo = append(m.DownloadInfo, utils.DownloadInfo{
			DownloadUrl:     strings.ReplaceAll(v.URL, constants.MinecraftLaunchMeta, utils.MirrorUtil{}.LaunchMetaURL()),
			DownloadStorage: storageManifestsPath(v.URL),
			Hash:            v.Sha1,
			Size:            0,
		})
	}
	return nil
}

func storageManifestsPath(url string) string {
	wd, _ := os.Getwd()
	p := wd + string(filepath.Separator) + "manifests" + string(filepath.Separator) + strings.ReplaceAll(url, constants.MinecraftLaunchMeta, "")
	if runtime.GOOS == "windows" {
		return filepath.FromSlash(p)
	}
	return p
}
