// Package array provides both threadsafe and non-threadsafe implementations of
// a generic dynamic array. In the threadsafe array, safety encompasses all
// operations on one array. Operations on multiple arrays are consistent in that
// the elements of each array used was valid at exactly one point in time
// between the start and the end of the operation.
package array

import "github.com/khezen/struct/collection"

// Interface is describing a Set. Sets are an unordered, unique list of values.
type Interface interface {
	collection.Interface
	Get(i int) (interface{}, error)
	Insert(i int, item ...interface{}) error
	RemoveAt(i int) (interface{}, error)
	ReplaceAt(i int, substitute interface{}) (interface{}, error)
	IndexOf(interface{}) (int, error)
	Swap(i, j int) error

	SubArray(i, j int) (Interface, error)
	Copy() Interface
}
