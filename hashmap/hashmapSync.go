package hashmap

import (
	"sync"
)

type hashmapSync struct {
	hashmap
	l sync.RWMutex
}

// NewSync creates a new thread safe hashmap
func NewSync(pairs ...interface{}) Interface {
	return &hashmapSync{
		*New(pairs...).(*hashmap),
		sync.RWMutex{},
	}
}

func (h *hashmapSync) Get(k interface{}) (interface{}, error) {
	h.l.RLock()
	defer h.l.RUnlock()
	return h.hashmap.Get(k)
}

func (h *hashmapSync) Put(k, v interface{}) {
	h.l.Lock()
	defer h.l.Unlock()
	h.hashmap.Put(k, v)
}

func (h *hashmapSync) Remove(keys ...interface{}) {
	h.l.Lock()
	defer h.l.Unlock()
	h.hashmap.Remove(keys...)
}

func (h *hashmapSync) Has(keys ...interface{}) bool {
	h.l.RLock()
	defer h.l.RUnlock()
	return h.hashmap.Has(keys...)
}

func (h *hashmapSync) HasValue(values ...interface{}) bool {
	h.l.RLock()
	defer h.l.RUnlock()
	return h.hashmap.HasValue(values...)
}

func (h *hashmapSync) KeyOf(value interface{}) (interface{}, error) {
	h.l.RLock()
	defer h.l.RUnlock()
	return h.hashmap.KeyOf(value)
}

func (h *hashmapSync) Each(f func(k, v interface{}) bool) {
	h.l.RLock()
	defer h.l.RUnlock()
	h.hashmap.Each(f)
}

func (h *hashmapSync) Len() int {
	h.l.RLock()
	defer h.l.RUnlock()
	return h.hashmap.Len()
}

func (h *hashmapSync) Clear() {
	h.l.Lock()
	defer h.l.Unlock()
	h.hashmap.Clear()
}

func (h *hashmapSync) IsEmpty() bool {
	h.l.RLock()
	defer h.l.RUnlock()
	return h.hashmap.IsEmpty()
}

func (h *hashmapSync) IsEqual(t Interface) bool {
	h.l.RLock()
	defer h.l.RUnlock()
	return h.hashmap.IsEqual(t)
}

func (h *hashmapSync) String() string {
	h.l.RLock()
	defer h.l.RUnlock()
	return h.hashmap.String()
}

func (h *hashmapSync) Keys() []interface{} {
	h.l.RLock()
	defer h.l.RUnlock()
	return h.hashmap.Keys()
}

func (h *hashmapSync) Values() []interface{} {
	h.l.RLock()
	defer h.l.RUnlock()
	return h.hashmap.Values()
}

func (h *hashmapSync) Copy() Interface {
	h.l.RLock()
	defer h.l.RUnlock()
	cpy := NewSync()
	h.Each(func(k, v interface{}) bool {
		cpy.Put(k, v)
		return true
	})
	return cpy
}
