package oset

import (
	"github.com/khezen/struct/array"
	"github.com/khezen/struct/collection"
	"github.com/khezen/struct/set"
	"sync"
)

type osetTS struct {
	oset
	l sync.RWMutex
}

// NewTS creates a thread safe ordered set
func NewTS(items ...interface{}) Interface {
	return &osetTS{
		*New(items...).(*oset),
		sync.RWMutex{},
	}
}

func (s *osetTS) Get(i int) (interface{}, error) {
	s.l.RLock()
	defer s.l.RUnlock()
	return s.oset.Get(i)
}

// Add includes the specified items (one or more) to the oset. The underlying
// osetTS. If passed nothing it silently returns.
func (s *osetTS) Add(items ...interface{}) {
	if len(items) > 0 {
		s.l.Lock()
		defer s.l.Unlock()
		s.oset.Add(items...)
	}
}

func (s *osetTS) Insert(i int, items ...interface{}) error {
	s.l.Lock()
	defer s.l.Unlock()
	return s.oset.Insert(i, items...)
}

// Remove deletes the specified items from the oset.  The underlying osetTS s is
// modified. If passed nothing it silently returns.
func (s *osetTS) Remove(items ...interface{}) {
	if len(items) > 0 {
		s.l.Lock()
		defer s.l.Unlock()
		s.oset.Remove(items...)
	}
}

func (s *osetTS) RemoveAt(i int) (interface{}, error) {
	s.l.Lock()
	defer s.l.Unlock()
	return s.oset.RemoveAt(i)
}

func (s *osetTS) Replace(toBeReplaced, substitute interface{}) {
	s.l.Lock()
	defer s.l.Unlock()
	s.oset.Replace(toBeReplaced, substitute)
}

func (s *osetTS) ReplaceAt(i int, substitute interface{}) (interface{}, error) {
	s.l.Lock()
	defer s.l.Unlock()
	return s.oset.ReplaceAt(i, substitute)
}

func (s *osetTS) IndexOf(item interface{}) (int, error) {
	s.l.RLock()
	defer s.l.RUnlock()
	return s.oset.IndexOf(item)
}

func (s *osetTS) Swap(i, j int) {
	s.l.Lock()
	defer s.l.Unlock()
	s.oset.Swap(i, j)
}

// Has looks for the existence of items passed. It returns false if nothing is
// passed. For multiple items it returns true only if all of  the items exist.
func (s *osetTS) Has(items ...interface{}) bool {
	switch len(items) {
	case 0:
		return true
	default:
		s.l.RLock()
		defer s.l.RUnlock()
		return s.oset.Has(items...)
	}

}

// Each traverses the items in the osetTS, calling the provided function for each
// oset member. Traversal will continue until all items in the osetTS have been
// visited, or if the closure returns false.
func (s *osetTS) Each(f func(item interface{}) bool) {
	s.l.RLock()
	defer s.l.RUnlock()
	s.oset.Each(f)
}

// Len returns the number of items in a oset.
func (s *osetTS) Len() int {
	s.l.RLock()
	defer s.l.RUnlock()
	return s.oset.Len()
}

// Clear removes all items from the oset.
func (s *osetTS) Clear() {
	s.l.Lock()
	defer s.l.Unlock()
	s.oset.Clear()
}

func (s *osetTS) IsEmpty() bool {
	s.l.RLock()
	defer s.l.RUnlock()
	return s.oset.IsEmpty()
}

func (s *osetTS) IsEqual(t collection.Interface) bool {
	s.l.RLock()
	defer s.l.RUnlock()
	return s.oset.IsEqual(t)
}

func (s *osetTS) IsSubset(t Interface) bool {
	s.l.RLock()
	defer s.l.RUnlock()
	return s.oset.IsSubset(t)
}

func (s *osetTS) IsSuperset(t Interface) bool {
	s.l.RLock()
	defer s.l.RUnlock()
	return s.oset.IsSuperset(t)
}

// Merge is like Union, however it modifies the current oset it's applied on
// with the given t oset.
func (s *osetTS) Merge(t collection.Interface) {
	if !t.IsEmpty() {
		s.l.Lock()
		defer s.l.Unlock()
		s.oset.Merge(t)
	}
}

func (s *osetTS) Separate(t collection.Interface) {
	if !t.IsEmpty() {
		s.l.Lock()
		defer s.l.Unlock()
		s.oset.Separate(t)
	}
}

func (s *osetTS) Retain(t collection.Interface) {
	s.l.Lock()
	defer s.l.Unlock()
	s.oset.Retain(t)
}

func (s *osetTS) SubArray(i, j int) (array.Interface, error) {
	s.l.RLock()
	defer s.l.RUnlock()
	arr, err := s.oset.SubArray(i, j)
	if err != nil {
		return nil, err
	}
	return array.NewTS(arr.Slice()...), nil
}

func (s *osetTS) Subset(i, j int) (Interface, error) {
	s.l.RLock()
	defer s.l.RUnlock()
	os, err := s.oset.Subset(i, j)
	if err != nil {
		return nil, err
	}
	return &osetTS{
		*os.(*oset),
		sync.RWMutex{},
	}, nil
}

func (s *osetTS) String() string {
	s.l.RLock()
	defer s.l.RUnlock()
	return s.oset.String()
}

// Slice returns a slice of all items. There is also StringSlice() and
// IntSlice() methods for returning slices of type string or int.
func (s *osetTS) Slice() []interface{} {
	s.l.RLock()
	defer s.l.RUnlock()
	return s.oset.Slice()
}

func (s *osetTS) CopyOset() Interface {
	return NewTS(s.Slice()...)
}

// Copy returns a new osetTS with a copy of s.
func (s *osetTS) CopyArr() array.Interface {
	s.l.RLock()
	defer s.l.RUnlock()
	return NewTS(s.oset.Slice()...)
}

func (s *osetTS) CopySet() set.Interface {
	s.l.RLock()
	defer s.l.RUnlock()
	return set.NewTS(s.Slice()...)
}

func (s *osetTS) CopyCollection() collection.Interface {
	return s.CopyArr()
}
