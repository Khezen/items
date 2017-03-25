// Package set provides both threadsafe and non-threadsafe implementations of
// a generic set data structure. In the threadsafe set, safety encompasses all
// operations on one set. Operations on multiple sets are consistent in that
// the elements of each set used was valid at exactly one point in time
// between the start and the end of the operation.
package set

// Interface is describing a set. sets are an unordered, unique Slice of values.
type Interface interface {
	Add(items ...interface{})
	Remove(items ...interface{})
	Pop() interface{}
	Has(items ...interface{}) bool
	Each(func(interface{}) bool)

	Len() int
	Clear()
	IsEmpty() bool
	IsEqual(s Interface) bool
	IsSubset(s Interface) bool
	IsSuperset(s Interface) bool
	Merge(s Interface)
	Separate(s Interface)
	Retain(s Interface)

	String() string
	Slice() []interface{}
	Copy() Interface
}

// helpful to not write everywhere struct{}{}
var keyExists = struct{}{}

// Union is the merger of multiple sets. It returns a new set with all the
// elements present in all the sets that are passed.
//
// The dynamic type of the returned set is determined by the first passed set's
// implementation of the New() method.
func Union(set1, set2 Interface, sets ...Interface) Interface {
	f := set1.Copy()
	set2.Each(func(item interface{}) bool {
		f.Add(item)
		return true
	})
	for _, set := range sets {
		set.Each(func(item interface{}) bool {
			f.Add(item)
			return true
		})
	}

	return f
}

// Difference returns a new set which contains items which are in in the first
// set but not in the others. Unlike the Difference() method you can use this
// function separately with multiple sets.
func Difference(set1, set2 Interface, sets ...Interface) Interface {
	s := set1.Copy()
	s.Separate(set2)
	for _, set := range sets {
		s.Separate(set) // seperate is thread safe
	}
	return s
}

// Intersection returns a new set which contains items that only exist in all given sets.
func Intersection(set1, set2 Interface, sets ...Interface) Interface {
	result := New(set1.Slice()...)
	result.Retain(set2)
	for _, set := range sets {
		result.Retain(set)
	}
	return result
}

// Exclusion returns a new set which s is the difference of items which are in
// one of either, but not in both.
func Exclusion(set1 Interface, set2 Interface, sets ...Interface) Interface {
	intersection := Intersection(set1, set2, sets...)
	exclusion := Union(set1, set2, sets...)
	exclusion.Separate(intersection)
	return exclusion
}
