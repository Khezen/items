// Package set provides both threadsafe and non-threadsafe implementations of
// a generic set data structure. In the threadsafe set, safety encompasses all
// operations on one set. Operations on multiple sets are consistent in that
// the elements of each set used was valid at exactly one point in time
// between the start and the end of the operation.
package set

import (
	"github.com/khezen/struct/collection"
)

// Interface is describing a set. sets are an unordered, unique Slice of values.
type Interface interface {
	collection.Interface
	Pop() interface{}
	IsSubset(s Interface) bool
	IsSuperset(s Interface) bool
	Copy() Interface
}

// helpful to not write everywhere struct{}{}
var keyExists = struct{}{}
