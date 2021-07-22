package providers

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/FabricMCCN/minecraft-mirror-software/constants"
	"github.com/FabricMCCN/minecraft-mirror-software/jsonstrs"
	"github.com/FabricMCCN/minecraft-mirror-software/logger"
	"github.com/FabricMCCN/minecraft-mirror-software/utils"
)

var rOnce sync.Once
var rIns *ResourceListProvider

type ResourceListProvider struct {
	DownloadInfo []utils.DownloadInfo
}

func NewResourceListProvider() *ResourceListProvider {
	rOnce.Do(func() { rIns = &ResourceListProvider{} })
	return rIns
}

func (r *ResourceListProvider) GetDownloadList() []utils.DownloadInfo { return r.DownloadInfo }

func (r *ResourceListProvider) Parse(b []byte) error {
	assetIndex := &jsonstrs.Resource{}
	if err := json.Unmarshal(b, assetIndex); err != nil {
		logger.Error(err)
		return err
	}

	for _, v := range assetIndex.Objects {
		r.DownloadInfo = append(r.DownloadInfo, utils.DownloadInfo{
			DownloadUrl:     constants.MinecraftAssets + v.Hash[0:2] + "/" + v.Hash,
			DownloadStorage: storageAssetsPath(v.Hash),
		})
	}

	return nil
}

func storageAssetsPath(hash string) string {
	wd, _ := os.Getwd()
	p := wd + string(filepath.Separator) + "assets" + string(filepath.Separator) + hash[0:2] + string(filepath.Separator) + hash
	if runtime.GOOS == "windows" {
		return filepath.FromSlash(p)
	}
	return p
}
