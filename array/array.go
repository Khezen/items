// Package array provides both threadsafe and non-threadsafe implementations of
// a generic dynamic array. In the threadsafe array, safety encompasses all
// operations on one array. Operations on multiple arrays are consistent in that
// the elements of each array used was valid at exactly one point in time
// between the start and the end of the operation.
package array

// Interface is describing a Set. Sets are an unordered, unique list of values.
type Interface interface {
	Get(i int) (interface{}, error)
	Add(items ...interface{})
	Insert(i int, item ...interface{}) error
	Remove(items ...interface{})
	RemoveAt(i int) (interface{}, error)
	Replace(toBeReplaced, substitute interface{})
	ReplaceAt(i int, substitute interface{}) (interface{}, error)
	IndexOf(interface{}) (int, error)
	Swap(i, j int) error
	Has(items ...interface{}) bool
	Each(func(interface{}) bool)

	Len() int
	Clear()
	IsEmpty() bool
	IsEqual(Interface) bool

	Merge(a Interface)
	Separate(a Interface)
	Retain(a Interface)

	String() string
	Slice() []interface{}
	Copy() Interface
}

// Union is the merger of multiple arrays. It returns a new array with all the
// elements present in all the arrays that are passed.
//
// The dynamic type of the returned array is determined by the first passed array's
// implementation of the New() method.
func Union(arrays ...Interface) Interface {
	if len(arrays) == 0 {
		return nil
	}

	u := arrays[0].Copy()
	arrays = arrays[1:]
	for _, array := range arrays {
		array.Each(func(item interface{}) bool {
			if !u.Has(item) {
				u.Add(item)
			}
			return true
		})
	}
	return u
}

// Difference returns a new array which contains items which are in in the first
// array but not in the others. Unlike the Difference() method you can use this
// function separately with multiple arrays.
func Difference(arrays ...Interface) Interface {
	if len(arrays) == 0 {
		return nil
	}

	s := arrays[0].Copy()
	arrays = arrays[1:]
	for _, array := range arrays {
		s.Separate(array) // seperate is thread safe
	}
	return s
}

// Intersection returns a new array which contains items that only exist in all given arrays.
func Intersection(arrays ...Interface) Interface {
	if len(arrays) == 0 {
		return nil
	}

	result := Union(arrays...)
	for i := 0; i < result.Len(); i++ {
		removed := false
		item, _ := result.Get(i)
		for _, array := range arrays {
			if !removed && !array.Has(item) {
				result.RemoveAt(i)
				removed = true
				break
			}
		}
		if removed {
			i--
		}
	}
	return result
}

// Exclusion returns a new array which s is the difference of items which are in
// one of either, but not in both.
func Exclusion(arrays ...Interface) Interface {
	length := len(arrays)
	if length == 0 {
		return nil
	}
	if length == 1 {
		return arrays[0]
	}
	intersections := make([]Interface, 0, length)
	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			if j != i {
				intersections = append(intersections, Intersection(arrays[i], arrays[j]))
			}
		}
	}
	intersection := Union(intersections...)
	exclusion := Union(arrays...)
	exclusion.Separate(intersection)
	return exclusion
}
