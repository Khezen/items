package oset

import (
	"sync"

	"github.com/khezen/struct/array"
	"github.com/khezen/struct/collection"
	"github.com/khezen/struct/set"
)

type osetSync struct {
	oset
	l sync.RWMutex
}

// NewSync creates a thread safe ordered set
func NewSync(items ...interface{}) Interface {
	return &osetSync{
		*New(items...).(*oset),
		sync.RWMutex{},
	}
}

func (s *osetSync) Get(i int) interface{} {
	s.l.RLock()
	defer s.l.RUnlock()
	return s.oset.Get(i)
}

// Add includes the specified items (one or more) to the oset. The underlying
// osetSync. If passed nothing it silently returns.
func (s *osetSync) Add(items ...interface{}) {
	if len(items) > 0 {
		s.l.Lock()
		defer s.l.Unlock()
		s.oset.Add(items...)
	}
}

func (s *osetSync) Insert(i int, items ...interface{}) {
	s.l.Lock()
	defer s.l.Unlock()
	s.oset.Insert(i, items...)
}

// Remove deletes the specified items from the oset.  The underlying osetSync s is
// modified. If passed nothing it silently returns.
func (s *osetSync) Remove(items ...interface{}) {
	if len(items) > 0 {
		s.l.Lock()
		defer s.l.Unlock()
		s.oset.Remove(items...)
	}
}

func (s *osetSync) RemoveAt(i int) interface{} {
	s.l.Lock()
	defer s.l.Unlock()
	return s.oset.RemoveAt(i)
}

func (s *osetSync) Replace(toBeReplaced, substitute interface{}) {
	s.l.Lock()
	defer s.l.Unlock()
	s.oset.Replace(toBeReplaced, substitute)
}

func (s *osetSync) ReplaceAt(i int, substitute interface{}) interface{} {
	s.l.Lock()
	defer s.l.Unlock()
	return s.oset.ReplaceAt(i, substitute)
}

func (s *osetSync) IndexOf(item interface{}) (int, error) {
	s.l.RLock()
	defer s.l.RUnlock()
	return s.oset.IndexOf(item)
}

func (s *osetSync) Swap(i, j int) {
	s.l.Lock()
	defer s.l.Unlock()
	s.oset.Swap(i, j)
}

// Has looks for the existence of items passed. It returns false if nothing is
// passed. For multiple items it returns true only if all of  the items exist.
func (s *osetSync) Has(items ...interface{}) bool {
	switch len(items) {
	case 0:
		return true
	default:
		s.l.RLock()
		defer s.l.RUnlock()
		return s.oset.Has(items...)
	}

}

// Each traverses the items in the osetSync, calling the provided function for each
// oset member. Traversal will continue until all items in the osetSync have been
// visited, or if the closure returns false.
func (s *osetSync) Each(f func(item interface{}) bool) {
	s.l.RLock()
	defer s.l.RUnlock()
	s.oset.Each(f)
}

// Len returns the number of items in a oset.
func (s *osetSync) Len() int {
	s.l.RLock()
	defer s.l.RUnlock()
	return s.oset.Len()
}

// Clear removes all items from the oset.
func (s *osetSync) Clear() {
	s.l.Lock()
	defer s.l.Unlock()
	s.oset.Clear()
}

func (s *osetSync) IsEmpty() bool {
	s.l.RLock()
	defer s.l.RUnlock()
	return s.oset.IsEmpty()
}

func (s *osetSync) IsEqual(t collection.Interface) bool {
	s.l.RLock()
	defer s.l.RUnlock()
	return s.oset.IsEqual(t)
}

func (s *osetSync) IsSubset(t Interface) bool {
	s.l.RLock()
	defer s.l.RUnlock()
	return s.oset.IsSubset(t)
}

func (s *osetSync) IsSuperset(t Interface) bool {
	s.l.RLock()
	defer s.l.RUnlock()
	return s.oset.IsSuperset(t)
}

// Merge is like Union, however it modifies the current oset it's applied on
// with the given t oset.
func (s *osetSync) Merge(t collection.Interface) {
	if !t.IsEmpty() {
		s.l.Lock()
		defer s.l.Unlock()
		s.oset.Merge(t)
	}
}

func (s *osetSync) Separate(t collection.Interface) {
	if !t.IsEmpty() {
		s.l.Lock()
		defer s.l.Unlock()
		s.oset.Separate(t)
	}
}

func (s *osetSync) Retain(t collection.Interface) {
	s.l.Lock()
	defer s.l.Unlock()
	s.oset.Retain(t)
}

func (s *osetSync) SubArray(i, j int) array.Interface {
	s.l.RLock()
	defer s.l.RUnlock()
	arr := s.oset.SubArray(i, j)
	return array.NewSync(arr.Slice()...)
}

func (s *osetSync) Subset(i, j int) Interface {
	s.l.RLock()
	defer s.l.RUnlock()
	os := s.oset.Subset(i, j)
	return &osetSync{
		*os.(*oset),
		sync.RWMutex{},
	}
}

func (s *osetSync) String() string {
	s.l.RLock()
	defer s.l.RUnlock()
	return s.oset.String()
}

// Slice returns a slice of all items. There is also StringSlice() and
// IntSlice() methods for returning slices of type string or int.
func (s *osetSync) Slice() []interface{} {
	s.l.RLock()
	defer s.l.RUnlock()
	return s.oset.Slice()
}

func (s *osetSync) CopyOset() Interface {
	return NewSync(s.Slice()...)
}

// Copy returns a new osetSync with a copy of s.
func (s *osetSync) CopyArr() array.Interface {
	s.l.RLock()
	defer s.l.RUnlock()
	return NewSync(s.oset.Slice()...)
}

func (s *osetSync) CopySet() set.Interface {
	s.l.RLock()
	defer s.l.RUnlock()
	return set.NewSync(s.Slice()...)
}

func (s *osetSync) CopyCollection() collection.Interface {
	return s.CopyArr()
}
