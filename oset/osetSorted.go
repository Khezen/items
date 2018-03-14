package oset

import (
	"sort"
)

type osetSort struct {
	oset
	less func(slice []interface{}, i, j int) bool
}

// NewSorted creates an oordered set that expose Sort method
func NewSorted(less func(slice []interface{}, i, j int) bool, items ...interface{}) Sorted {
	return &osetSort{
		*(New(items...).(*oset)),
		less,
	}
}

func (a *osetSort) Less(i, j int) bool {
	return a.less(a.Slice(), i, j)
}

func (a *osetSort) Sort() {
	sort.Sort(a)
}
