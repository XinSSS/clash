package nat

import (
	"sync"

	C "github.com/XinSSS/clash/constant"
)

type Table struct {
	mapping sync.Map
}

func (t *Table) Set(key string, pc C.PacketConn) {
	t.mapping.Store(key, pc)
}

func (t *Table) Get(key string) C.PacketConn {
	item, exist := t.mapping.Load(key)
	if !exist {
		return nil
	}
	return item.(C.PacketConn)
}

func (t *Table) GetOrCreateLock(key string) (*sync.WaitGroup, bool) {
	item, loaded := t.mapping.LoadOrStore(key, &sync.WaitGroup{})
	return item.(*sync.WaitGroup), loaded
}

func (t *Table) Delete(key string) {
	t.mapping.Delete(key)
}

// New return *Cache
func New() *Table {
	return &Table{}
}
