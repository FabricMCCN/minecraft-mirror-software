package providers

import "github.com/FabricMCCN/minecraft-mirror-software/utils"

type Providers interface {
	GetDownloadList() []utils.DownloadInfo
	Parse(b []byte) error
}
