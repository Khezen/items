package array

import (
	"fmt"
	"github.com/khezen/check"
	"github.com/khezen/items/collection"
	"reflect"
	"strings"
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

func (a *array) Get(i int) (interface{}, error) {
	err := check.Index(i, a.Len())
	if err != nil {
		return nil, err
	}
	return a.s[i], nil
}

func (a *array) Add(items ...interface{}) {
	if len(items) > 0 {
		a.s = append(a.s, items...)
	}
}

func (a *array) Insert(i int, items ...interface{}) error {
	err := check.Index(i, a.Len())
	if err != nil {
		return err
	}
	if len(items) > 0 {
		a.s = append(a.s[:i], append(items, a.s[i:]...)...)
	}
	return nil
}

func (a *array) Remove(items ...interface{}) {
	for _, item := range items {
		i, err := a.IndexOf(item)
		if err == nil {
			a.RemoveAt(i)
		}
	}
}

func (a *array) RemoveAt(i int) (interface{}, error) {
	err := check.Index(i, a.Len())
	if err != nil {
		return nil, err
	}
	item := a.s[i]
	a.s = append(a.s[:i], a.s[i+1:]...)
	return item, nil
}

func (a *array) Replace(toBeReplaced, substitute interface{}) {
	i, err := a.IndexOf(toBeReplaced)
	if err == nil {
		a.ReplaceAt(i, substitute)
	}
}

func (a *array) ReplaceAt(i int, substitute interface{}) (interface{}, error) {
	err := check.Index(i, a.Len())
	if err != nil {
		return nil, err
	}
	item := a.s[i]
	a.s[i] = substitute
	return item, nil
}

func (a *array) IndexOf(item interface{}) (int, error) {
	for i, current := range a.s {
		if item == current {
			return i, nil
		}
	}
	return -1, fmt.Errorf("not found")
}

func (a *array) Swap(i, j int) error {
	itemi, err := a.Get(i)
	if err != nil {
		return err
	}
	itemj, err := a.Get(j)
	if err != nil {
		return err
	}
	a.ReplaceAt(i, itemj)
	a.ReplaceAt(j, itemi)
	return nil
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
	a.Remove(t.Slice()...)
}

func (a *array) Retain(t collection.Interface) {
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

// Copy returns a new Set with a copy of s.
func (a *array) Copy() Interface {
	return New(a.s...)
}

func (a *array) CopyCollection() collection.Interface {
	return a.Copy()
}
