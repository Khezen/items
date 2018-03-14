package array

import (
	"sync"

	"github.com/khezen/struct/collection"
)

// arraySync defines a thread safe array data structure.
type arraySync struct {
	array
	l sync.RWMutex // we name it because we don't want to expose it
}

// NewSync creates a thread safe array
func NewSync(items ...interface{}) Interface {
	return &arraySync{
		*New(items...).(*array),
		sync.RWMutex{},
	}
}

func (a *arraySync) Get(i int) (interface{}, error) {
	a.l.RLock()
	defer a.l.RUnlock()
	return a.array.Get(i)
}

// Add includes the specified items (one or more) to the array. The underlying
// arraySync s is modified. If passed nothing it silently returns.
func (a *arraySync) Add(items ...interface{}) {
	if len(items) > 0 {
		a.l.Lock()
		defer a.l.Unlock()
		a.array.Add(items...)
	}
}

func (a *arraySync) Insert(i int, items ...interface{}) error {
	a.l.Lock()
	defer a.l.Unlock()
	return a.array.Insert(i, items...)
}

// Remove deletes the specified items from the array.  The underlying arraySync s is
// modified. If passed nothing it silently returns.
func (a *arraySync) Remove(items ...interface{}) {
	if len(items) > 0 {
		a.l.Lock()
		defer a.l.Unlock()
		a.array.Remove(items...)
	}
}

func (a *arraySync) RemoveAt(i int) (interface{}, error) {
	a.l.Lock()
	defer a.l.Unlock()
	return a.array.RemoveAt(i)
}

func (a *arraySync) Replace(toBeReplaced, substitute interface{}) {
	a.l.Lock()
	defer a.l.Unlock()
	a.array.Replace(toBeReplaced, substitute)
}

func (a *arraySync) ReplaceAt(i int, substitute interface{}) (interface{}, error) {
	a.l.Lock()
	defer a.l.Unlock()
	return a.array.ReplaceAt(i, substitute)
}

func (a *arraySync) IndexOf(item interface{}) (int, error) {
	a.l.RLock()
	defer a.l.RUnlock()
	return a.array.IndexOf(item)
}

func (a *arraySync) Swap(i, j int) {
	a.l.Lock()
	defer a.l.Unlock()
	a.array.Swap(i, j)
}

// Has looks for the existence of items passed. It returns false if nothing is
// passed. For multiple items it returns true only if all of  the items exist.
func (a *arraySync) Has(items ...interface{}) bool {
	switch len(items) {
	case 0:
		return true
	default:
		a.l.RLock()
		defer a.l.RUnlock()
		return a.array.Has(items...)
	}

}

// Each traverses the items in the arraySync, calling the provided function for each
// array member. Traversal will continue until all items in the arraySync have been
// visited, or if the closure returns false.
func (a *arraySync) Each(f func(item interface{}) bool) {
	a.l.RLock()
	defer a.l.RUnlock()
	a.array.Each(f)
}

// Len returns the number of items in a array.
func (a *arraySync) Len() int {
	a.l.RLock()
	defer a.l.RUnlock()
	return a.array.Len()
}

// Clear removes all items from the array.
func (a *arraySync) Clear() {
	a.l.Lock()
	defer a.l.Unlock()
	a.array.Clear()
}

func (a *arraySync) IsEmpty() bool {
	a.l.RLock()
	defer a.l.RUnlock()
	return a.array.IsEmpty()
}

func (a *arraySync) IsEqual(t collection.Interface) bool {
	a.l.RLock()
	defer a.l.RUnlock()
	return a.array.IsEqual(t)
}

// Merge is like Union, however it modifies the current array it's applied on
// with the given t array.
func (a *arraySync) Merge(t collection.Interface) {
	if !t.IsEmpty() {
		a.l.Lock()
		defer a.l.Unlock()
		a.array.Merge(t)
	}
}

func (a *arraySync) Separate(t collection.Interface) {
	if !t.IsEmpty() {
		a.l.Lock()
		defer a.l.Unlock()
		a.array.Separate(t)
	}
}

func (a *arraySync) Retain(t collection.Interface) {
	a.l.Lock()
	defer a.l.Unlock()
	a.array.Retain(t)
}

func (a *arraySync) String() string {
	a.l.RLock()
	defer a.l.RUnlock()
	return a.array.String()
}

// Slice returns a slice of all items. There is also StringSlice() and
// IntSlice() methods for returning slices of type string or int.
func (a *arraySync) Slice() []interface{} {
	a.l.RLock()
	defer a.l.RUnlock()
	return a.array.Slice()
}

// Copy returns a new arraySync with a copy of s.
func (a *arraySync) CopyArr() Interface {
	a.l.RLock()
	defer a.l.RUnlock()
	return NewSync(a.s...)
}

func (a *arraySync) SubArray(i, j int) (Interface, error) {
	a.l.RLock()
	defer a.l.RUnlock()
	arr, err := a.array.SubArray(i, j)
	if err != nil {
		return nil, err
	}
	return &arraySync{
		*arr.(*array),
		sync.RWMutex{},
	}, nil
}

func (a *arraySync) CopyCollection() collection.Interface {
	return a.CopyArr()
}
