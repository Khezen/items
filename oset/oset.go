package oset

import (
	"github.com/khezen/struct/array"
)

// Interface describe an ordered set
type Interface interface {
	array.Interface
	IsSubset(s Interface) bool
	IsSuperset(s Interface) bool
	CopyOset() Interface
	SubSet(i, j int) (Interface, error)
}
