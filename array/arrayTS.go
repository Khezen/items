package array

import (
	"github.com/khezen/struct/collection"
	"sync"
)

// arrayTS defines a thread safe array data structure.
type arrayTS struct {
	array
	l sync.RWMutex // we name it because we don't want to expose it
}

// NewTS creates a thread safe array
func NewTS(items ...interface{}) Interface {
	return &arrayTS{
		*New(items...).(*array),
		sync.RWMutex{},
	}
}

func (a *arrayTS) Get(i int) (interface{}, error) {
	a.l.RLock()
	defer a.l.RUnlock()
	return a.array.Get(i)
}

// Add includes the specified items (one or more) to the array. The underlying
// arrayTS s is modified. If passed nothing it silently returns.
func (a *arrayTS) Add(items ...interface{}) {
	if len(items) > 0 {
		a.l.Lock()
		defer a.l.Unlock()
		a.array.Add(items...)
	}
}

func (a *arrayTS) Insert(i int, items ...interface{}) error {
	a.l.Lock()
	defer a.l.Unlock()
	return a.array.Insert(i, items...)
}

// Remove deletes the specified items from the array.  The underlying arrayTS s is
// modified. If passed nothing it silently returns.
func (a *arrayTS) Remove(items ...interface{}) {
	if len(items) > 0 {
		a.l.Lock()
		defer a.l.Unlock()
		a.array.Remove(items...)
	}
}

func (a *arrayTS) RemoveAt(i int) (interface{}, error) {
	a.l.Lock()
	defer a.l.Unlock()
	return a.array.RemoveAt(i)
}

func (a *arrayTS) Replace(toBeReplaced, substitute interface{}) {
	a.l.Lock()
	defer a.l.Unlock()
	a.array.Replace(toBeReplaced, substitute)
}

func (a *arrayTS) ReplaceAt(i int, substitute interface{}) (interface{}, error) {
	a.l.Lock()
	defer a.l.Unlock()
	return a.array.ReplaceAt(i, substitute)
}

func (a *arrayTS) IndexOf(item interface{}) (int, error) {
	a.l.RLock()
	defer a.l.RUnlock()
	return a.array.IndexOf(item)
}

func (a *arrayTS) Swap(i, j int) {
	a.l.Lock()
	defer a.l.Unlock()
	a.array.Swap(i, j)
}

// Has looks for the existence of items passed. It returns false if nothing is
// passed. For multiple items it returns true only if all of  the items exist.
func (a *arrayTS) Has(items ...interface{}) bool {
	switch len(items) {
	case 0:
		return true
	default:
		a.l.RLock()
		defer a.l.RUnlock()
		return a.array.Has(items...)
	}

}

// Each traverses the items in the arrayTS, calling the provided function for each
// array member. Traversal will continue until all items in the arrayTS have been
// visited, or if the closure returns false.
func (a *arrayTS) Each(f func(item interface{}) bool) {
	a.l.RLock()
	defer a.l.RUnlock()
	a.array.Each(f)
}

// Len returns the number of items in a array.
func (a *arrayTS) Len() int {
	a.l.RLock()
	defer a.l.RUnlock()
	return a.array.Len()
}

// Clear removes all items from the array.
func (a *arrayTS) Clear() {
	a.l.Lock()
	defer a.l.Unlock()
	a.array.Clear()
}

func (a *arrayTS) IsEmpty() bool {
	a.l.RLock()
	defer a.l.RUnlock()
	return a.array.IsEmpty()
}

func (a *arrayTS) IsEqual(t collection.Interface) bool {
	a.l.RLock()
	defer a.l.RUnlock()
	return a.array.IsEqual(t)
}

// Merge is like Union, however it modifies the current array it's applied on
// with the given t array.
func (a *arrayTS) Merge(t collection.Interface) {
	if !t.IsEmpty() {
		a.l.Lock()
		defer a.l.Unlock()
		a.array.Merge(t)
	}
}

func (a *arrayTS) Separate(t collection.Interface) {
	if !t.IsEmpty() {
		a.l.Lock()
		defer a.l.Unlock()
		a.array.Separate(t)
	}
}

func (a *arrayTS) Retain(t collection.Interface) {
	a.l.Lock()
	defer a.l.Unlock()
	a.array.Retain(t)
}

func (a *arrayTS) String() string {
	a.l.RLock()
	defer a.l.RUnlock()
	return a.array.String()
}

// Slice returns a slice of all items. There is also StringSlice() and
// IntSlice() methods for returning slices of type string or int.
func (a *arrayTS) Slice() []interface{} {
	a.l.RLock()
	defer a.l.RUnlock()
	return a.array.Slice()
}

// Copy returns a new arrayTS with a copy of s.
func (a *arrayTS) CopyArr() Interface {
	a.l.RLock()
	defer a.l.RUnlock()
	return NewTS(a.s...)
}

func (a *arrayTS) SubArray(i, j int) (Interface, error) {
	a.l.RLock()
	defer a.l.RUnlock()
	arr, err := a.array.SubArray(i, j)
	if err != nil {
		return nil, err
	}
	return &arrayTS{
		*arr.(*array),
		sync.RWMutex{},
	}, nil
}

func (a *arrayTS) CopyCollection() collection.Interface {
	return a.CopyArr()
}
