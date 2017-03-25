package set

import (
	"github.com/khezen/struct/collection"
	"sync"
)

// setTS defines a thread safe set data structure.
type setTS struct {
	set
	l sync.RWMutex // we name it because we don't want to expose it
}

// NewTS creates and initialize a new thread safe set. It's accept a variable number of
// arguments to populate the initial set. If nothing passed a set with zero
// size is created.
func NewTS(items ...interface{}) Interface {
	return &setTS{
		*New(items...).(*set),
		sync.RWMutex{},
	}
}

// Add includes the specified items (one or more) to the set. The underlying
// set s is modified. If passed nothing it silently returns.
func (s *setTS) Add(items ...interface{}) {
	s.l.Lock()
	defer s.l.Unlock()
	s.set.Add(items...)
}

// Remove deletes the specified items from the set.  The underlying set s is
// modified. If passed nothing it silently returns.
func (s *setTS) Remove(items ...interface{}) {
	s.l.Lock()
	defer s.l.Unlock()
	s.set.Remove(items...)
}

// Pop  deletes and return an item from the set. The underlying set s is
// modified. If set is empty, nil is returned.
func (s *setTS) Pop() interface{} {
	s.l.Lock()
	defer s.l.Unlock()
	return s.set.Pop()
}

// Has looks for the existence of items passed. It returns false if nothing is
// passed. For multiple items it returns true only if all of  the items exist.
func (s *setTS) Has(items ...interface{}) bool {
	s.l.RLock()
	defer s.l.RUnlock()
	return s.set.Has(items...)
}

func (s *setTS) Replace(item, substitute interface{}) {
	s.l.Lock()
	defer s.l.Unlock()
	s.set.Replace(item, substitute)
}

// Len returns the number of items in a set.
func (s *setTS) Len() int {
	s.l.RLock()
	defer s.l.RUnlock()
	return s.set.Len()
}

// Clear removes all items from the set.
func (s *setTS) Clear() {
	s.l.Lock()
	defer s.l.Unlock()
	s.set.Clear()
}

// IsEqual test whether s and t are the same in size and have the same items.
func (s *setTS) IsEqual(t collection.Interface) bool {
	s.l.RLock()
	defer s.l.RUnlock()
	return s.set.IsEqual(t)
}

// IsSubset tests whether t is a subset of s.
func (s *setTS) IsSubset(t Interface) (subset bool) {
	s.l.RLock()
	defer s.l.RUnlock()
	return s.set.IsSubset(t)
}

// Each traverses the items in the set, calling the provided function for each
// set member. Traversal will continue until all items in the set have been
// visited, or if the closure returns false.
func (s *setTS) Each(f func(item interface{}) bool) {
	s.l.RLock()
	defer s.l.RUnlock()
	s.set.Each(f)
}

// Slice returns a slice of all items. There is also StringSlice() and
// IntSlice() methods for returning slices of type string or int.
func (s *setTS) Slice() []interface{} {
	s.l.RLock()
	defer s.l.RUnlock()
	return s.set.Slice()
}

// Copy returns a new set with a copy of s.
func (s *setTS) Copy() Interface {
	s.l.RLock()
	defer s.l.RUnlock()
	u := NewTS()
	for item := range s.m {
		u.Add(item)
	}
	return u
}

// Merge is like Union, however it modifies the current set it's applied on
// with the given t set.
func (s *setTS) Merge(t collection.Interface) {
	s.l.Lock()
	defer s.l.Unlock()
	s.set.Merge(t)
}

// Retain removes the set items not containing in t from set s.
func (s *setTS) Retain(t collection.Interface) {
	s.l.Lock()
	defer s.l.Unlock()
	s.set.Retain(t)
}

func (s *setTS) CopyCollection() collection.Interface {
	return s.Copy()
}
