package set

import (
	"fmt"
	"github.com/khezen/struct/collection"
	"strings"
)

// set defines a non-thread safe set data structure.
type set struct {
	m map[interface{}]struct{} // struct{} doesn't take up space
}

// New creates and initializes a new non-threadsafe set.
func New(items ...interface{}) Interface {
	s := &set{}
	s.m = make(map[interface{}]struct{})
	s.Add(items...)
	return s
}

// Add includes the specified items (one or more) to the set. The underlying
// set s is modified. If passed nothing it silently returns.
func (s *set) Add(items ...interface{}) {
	if len(items) > 0 {
		for _, item := range items {
			s.m[item] = keyExists
		}
	}
}

// Remove deletes the specified items from the set.  The underlying set s is
// modified. If passed nothing it silently returns.
func (s *set) Remove(items ...interface{}) {
	if len(items) > 0 {
		for _, item := range items {
			delete(s.m, item)
		}
	}
}

// Pop  deletes and return an item from the set. The underlying set s is
// modified. If set is empty, nil is returned.
func (s *set) Pop() interface{} {
	for item := range s.m {
		delete(s.m, item)
		return item
	}
	return nil
}

// Has looks for the existence of items passed. It returns false if nothing is
// passed. For multiple items it returns true only if all of  the items exist.
func (s *set) Has(items ...interface{}) bool {
	has := true
	// assume checked for empty item, which not exist
	if len(items) > 0 {
		for _, item := range items {
			if _, has = s.m[item]; !has {
				break
			}
		}
	}
	return has
}

func (s *set) Replace(item, substitute interface{}) {
	if _, ok := s.m[item]; ok {
		delete(s.m, item)
		s.m[substitute] = keyExists
	}
}

// Len returns the number of items in a set.
func (s *set) Len() int {
	return len(s.m)
}

// Clear removes all items from the set.
func (s *set) Clear() {
	s.m = make(map[interface{}]struct{})
}

// IsEmpty reports whether the set is empty.
func (s *set) IsEmpty() bool {
	return s.Len() == 0
}

// IsEqual test whether s and t are the same in size and have the same items.
func (s *set) IsEqual(t collection.Interface) bool {
	// Force locking only if given set is threadsafe.
	if conv, ok := t.(*setTS); ok {
		conv.l.RLock()
		defer conv.l.RUnlock()
	}

	// return false if they are no the same size
	if len(s.m) != t.Len() {
		return false
	}

	equal := true
	t.Each(func(item interface{}) bool {
		_, ok := s.m[item]
		equal = equal && ok
		return equal // if false, Each() will end
	})

	return equal
}

// IsSubset tests whether t is a subset of s.
func (s *set) IsSubset(t Interface) (subset bool) {
	subset = true

	t.Each(func(item interface{}) bool {
		_, subset = s.m[item]
		return subset
	})

	return
}

// IsSuperset tests whether t is a superset of s.
func (s *set) IsSuperset(t Interface) bool {
	return t.IsSubset(s)
}

// Each traverses the items in the set, calling the provided function for each
// set member. Traversal will continue until all items in the set have been
// visited, or if the closure returns false.
func (s *set) Each(f func(item interface{}) bool) {
	for item := range s.m {
		if !f(item) {
			break
		}
	}
}

// Copy returns a new set with a copy of s.
func (s *set) Copy() Interface {
	u := New()
	for item := range s.m {
		u.Add(item)
	}
	return u
}

// String returns a string representation of s
func (s *set) String() string {
	t := make([]string, 0, len(s.Slice()))
	for _, item := range s.Slice() {
		t = append(t, fmt.Sprintf("%v", item))
	}

	return fmt.Sprintf("[%s]", strings.Join(t, " "))
}

// Slice returns a slice of all items. There is also StringSlice() and
// IntSlice() methods for returning slices of type string or int.
func (s *set) Slice() []interface{} {
	Slice := make([]interface{}, 0, len(s.m))

	for item := range s.m {
		Slice = append(Slice, item)
	}

	return Slice
}

// Merge is like Union, however it modifies the current set it's applied on
// with the given t set.
func (s *set) Merge(t collection.Interface) {
	t.Each(func(item interface{}) bool {
		s.m[item] = keyExists
		return true
	})
}

// it's not the opposite of Merge.
// Separate removes the set items containing in t from set s.
func (s *set) Separate(t collection.Interface) {
	s.Remove(t.Slice()...)
}

// Retain removes the set items not containing in t from set s.
func (s *set) Retain(t collection.Interface) {
	items := make(map[interface{}]struct{})
	t.Each(func(item interface{}) bool {
		if s.Has(item) {
			items[item] = keyExists
		}
		return true
	})
	s.m = items
}

func (s *set) CopyCollection() collection.Interface {
	return s.Copy()
}
