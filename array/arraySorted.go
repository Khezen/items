package array

import (
	"sort"
)

// Sortable is the interface for sortable arrays
type Sortable interface {
	Sort()
	Less(i, j int) bool
}

type arraySort struct {
	array
	less func(i, j int) bool
}

// NewSortableArray creates an array that expose Sort method
func NewSortableArray(less func(i, j int) bool, items ...interface{}) Sortable {
	return &arraySort{
		*(New(items...).(*array)),
		less,
	}
}

func (a *arraySort) Less(i, j int) bool {
	return a.less(i, j)
}

func (a *arraySort) Sort() {
	sort.Sort(sort.Interface(a))
}
