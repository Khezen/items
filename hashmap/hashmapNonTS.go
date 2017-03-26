package hashmap

import (
	"fmt"
	"strings"
)

type hashmap struct {
	m map[interface{}]interface{}
}

// New creates a new hashmap
func New(pairs ...interface{}) Interface {
	h := &hashmap{
		m: make(map[interface{}]interface{}),
	}
	length := len(pairs)
	for i := 0; i < length-1; i += 2 {
		h.m[pairs[i]] = pairs[i+1]
	}
	return h
}

func (h *hashmap) Get(k interface{}) (interface{}, error) {
	v, ok := h.m[k]
	if !ok {
		return nil, fmt.Errorf("%v not found", k)
	}
	return v, nil
}

func (h *hashmap) Put(k, v interface{}) {
	h.m[k] = v
}

func (h *hashmap) Remove(keys ...interface{}) {
	for _, k := range keys {
		if _, ok := h.m[k]; ok {
			delete(h.m, k)
		}
	}
}

func (h *hashmap) Has(keys ...interface{}) bool {
	has := true
	for _, k := range keys {
		_, ok := h.m[k]
		has = has && ok
		if !has {
			return has
		}
	}
	return has
}

func (h *hashmap) HasValue(values ...interface{}) bool {
	has := true
	for _, value := range values {
		_, err := h.KeyOf(value)
		has = has && err == nil
		if !has {
			return has
		}
	}
	return has
}

func (h *hashmap) KeyOf(value interface{}) (interface{}, error) {
	for k, v := range h.m {
		if v == value {
			return k, nil
		}
	}
	return nil, fmt.Errorf("%v not found", value)
}

func (h *hashmap) Each(f func(k, v interface{}) bool) {
	for k, v := range h.m {
		if !f(k, v) {
			break
		}
	}
}

func (h *hashmap) Len() int {
	return len(h.m)
}

func (h *hashmap) Clear() {
	h.m = make(map[interface{}]interface{})
}

func (h *hashmap) IsEmpty() bool {
	return h.Len() == 0
}

func (h *hashmap) IsEqual(t Interface) bool {
	// Force locking only if given set is threadsafe.
	if conv, ok := t.(*hashmapTS); ok {
		conv.l.RLock()
		defer conv.l.RUnlock()
	}
	// return false if they are no the same size
	if sameLen := h.Len() == t.Len(); !sameLen {
		return false
	}
	equal := true
	t.Each(func(k, v interface{}) bool {
		value, err := h.Get(k)
		equal = equal && err != nil && value == v
		return equal // if false, Each() will end
	})
	return equal
}

func (h *hashmap) String() string {
	str := ""
	h.Each(func(k, v interface{}) bool {
		str = fmt.Sprintf("%v:%v ", k, v)
		return true
	})
	return fmt.Sprintf("[%v]", strings.Trim(str, " "))
}

func (h *hashmap) Keys() []interface{} {
	keys := make([]interface{}, 0, h.Len())
	h.Each(func(k, v interface{}) bool {
		keys = append(keys, k)
		return true
	})
	return keys
}

func (h *hashmap) Values() []interface{} {
	values := make([]interface{}, 0, h.Len())
	h.Each(func(k, v interface{}) bool {
		values = append(values, v)
		return true
	})
	return values
}

func (h *hashmap) Copy() Interface {
	cpy := New()
	h.Each(func(k, v interface{}) bool {
		cpy.Put(k, v)
		return true
	})
	return cpy
}
