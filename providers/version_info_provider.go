package providers

import "sync"

var vOnce sync.Once
var vIns *VersionListProvider

type VersionListProvider struct {
}

func NewVersionListProvider() *VersionListProvider {
	vOnce.Do(func() { vIns = &VersionListProvider{} })
	return vIns
}
