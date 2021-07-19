package providers

import "sync"

var rOnce sync.Once
var rIns *ResourceListProvider

type ResourceListProvider struct {
}

func NewResourceListProvider() *ResourceListProvider {
	rOnce.Do(func() { rIns = &ResourceListProvider{} })
	return rIns
}
