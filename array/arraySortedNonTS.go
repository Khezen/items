package array

import (
	"sort"
)

type arraySort struct {
	array
	less func(slice []interface{}, i, j int) bool
}

// NewSortableArray creates an array that expose Sort method
func NewSortableArray(less func(slice []interface{}, i, j int) bool, items ...interface{}) Sortable {
	return &arraySort{
		*(New(items...).(*array)),
		less,
	}
}

func (a *arraySort) Less(i, j int) bool {
	return a.less(a.Slice(), i, j)
}

func (a *arraySort) Sort() {
	sort.Sort(a)
}
