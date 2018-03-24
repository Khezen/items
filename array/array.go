package array

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/khezen/struct/collection"
)

// Provides a common array baseline for both threadsafe and non-ts arrays.
type array struct {
	s []interface{}
}

// New creates a non thread safe dynamic array
func New(items ...interface{}) Interface {
	size := len(items)
	a := &array{
		make([]interface{}, 0, size),
	}
	a.Add(items...)
	return a
}

func (a *array) Get(i int) interface{} {
	a.checkIndex(i)
	return a.s[i]
}

func (a *array) Add(items ...interface{}) {
	if len(items) > 0 {
		a.s = append(a.s, items...)
	}
}

func (a *array) Insert(i int, items ...interface{}) {
	a.checkIndex(i)
	if len(items) > 0 {
		a.s = append(a.s[:i], append(items, a.s[i:]...)...)
	}
}

func (a *array) Remove(items ...interface{}) {
	for _, item := range items {
		i, err := a.IndexOf(item)
		if err == nil {
			a.RemoveAt(i)
		}
	}
}

func (a *array) RemoveAt(i int) interface{} {
	a.checkIndex(i)
	item := a.s[i]
	length := len(a.s)
	copy(a.s[i:], a.s[i+1:])
	a.s[length-1] = nil
	a.s = a.s[:length-1]
	return item
}

func (a *array) Replace(toBeReplaced, substitute interface{}) {
	i, err := a.IndexOf(toBeReplaced)
	if err == nil {
		a.ReplaceAt(i, substitute)
	}
}

func (a *array) ReplaceAt(i int, substitute interface{}) interface{} {
	a.checkIndex(i)
	item := a.s[i]
	a.s[i] = substitute
	return item
}

func (a *array) IndexOf(item interface{}) (int, error) {
	for i, current := range a.s {
		if item == current {
			return i, nil
		}
	}
	return -1, ErrNotFound
}

func (a *array) Swap(i, j int) {
	itemi := a.Get(i)
	itemj := a.Get(j)
	a.ReplaceAt(i, itemj)
	a.ReplaceAt(j, itemi)
}

func (a *array) Has(items ...interface{}) bool {
	has := true
	for _, item := range items {
		_, err := a.IndexOf(item)
		has = has && err == nil
		if !has {
			return has
		}
	}
	return has
}

func (a *array) Each(f func(item interface{}) bool) {
	for _, item := range a.s {
		if !f(item) {
			break
		}
	}
}

// Len returns the number of items in a array.
func (a *array) Len() int {
	return len(a.s)
}

// Clear removes all items from the array.
func (a *array) Clear() {
	a.s = make([]interface{}, 0, 1)
}

// IsEmpty reports whether the Set is empty.
func (a *array) IsEmpty() bool {
	return a.Len() == 0
}

func (a *array) IsEqual(t collection.Interface) bool {
	// Force locking only if given set is threadsafe.
	if conv, ok := t.(*arraySync); ok {
		conv.l.RLock()
		defer conv.l.RUnlock()
	}
	if conv, ok := t.(*arraySortSync); ok {
		conv.l.RLock()
		defer conv.l.RUnlock()
	}
	length := a.Len()
	if length != t.Len() {
		return false
	}
	items := t.Slice()
	for i, item := range a.Slice() {
		compared := items[i]
		if reflect.TypeOf(item) != reflect.TypeOf(compared) {
			return false
		}
		if item != compared {
			return false
		}
	}
	return true
}

// Merge is like Union, however it modifies the current array it's applied on
// with the given t array.
func (a *array) Merge(t collection.Interface) {
	if conv, ok := t.(*arraySync); ok {
		conv.l.RLock()
		defer conv.l.RUnlock()
	}
	if conv, ok := t.(*arraySortSync); ok {
		conv.l.RLock()
		defer conv.l.RUnlock()
	}
	t.Each(func(item interface{}) bool {
		if !a.Has(item) {
			a.Add(item)
		}
		return true
	})
}

// it's not the opposite of Merge.
// Separate removes the array items containing in t from array s. Please aware that
func (a *array) Separate(t collection.Interface) {
	if conv, ok := t.(*arraySync); ok {
		conv.l.RLock()
		defer conv.l.RUnlock()
	}
	if conv, ok := t.(*arraySortSync); ok {
		conv.l.RLock()
		defer conv.l.RUnlock()
	}
	a.Remove(t.Slice()...)
}

func (a *array) Retain(t collection.Interface) {
	if conv, ok := t.(*arraySync); ok {
		conv.l.RLock()
		defer conv.l.RUnlock()
	}
	if conv, ok := t.(*arraySortSync); ok {
		conv.l.RLock()
		defer conv.l.RUnlock()
	}
	arr := make([]interface{}, 0, a.Len())
	a.Each(func(item interface{}) bool {
		if t.Has(item) {
			arr = append(arr, item)
		}
		return true
	})
	a.s = arr
}

// String returns a string representation of s
func (a *array) String() string {
	if a.IsEmpty() {
		return "[]"
	}
	t := make([]string, 0, a.Len())
	for _, item := range a.s {
		t = append(t, fmt.Sprintf("%v", item))
	}
	return fmt.Sprintf("[%s]", strings.Join(t, " "))
}

// Slice returns a slice of all items. There is also StringSlice() and
// IntSlice() methods for returning slices of type string or int.
func (a *array) Slice() []interface{} {
	return a.s
}

func (a *array) SubArray(i, j int) Interface {
	if i > j {
		panic(ErrBadSubsetBoudaries)
	}
	a.checkIndex(i)
	a.checkIndex(j)
	slice := a.Slice()
	result := New(slice...)
	result.Remove(slice[:i]...)
	result.Remove(slice[j+1:]...)
	return result
}

// Copy returns a new Set with a copy of s.
func (a *array) CopyArr() Interface {
	return New(a.s...)
}

func (a *array) CopyCollection() collection.Interface {
	return a.CopyArr()
}

func (a *array) checkIndex(i int) {
	if i < 0 || i >= a.Len() {
		panic(ErrIndexOutOfBounds)
	}
}
