package array

import (
	"sort"
)

// Sortable is the interface for sortable arrays
type arraySortTS struct {
	arrayTS
	less func(i, j int) bool
}

// NewSortableArrayTS creates an array that expose Sort method
func NewSortableArrayTS(less func(i, j int) bool, items ...interface{}) Sortable {
	return &arraySort{
		*(NewTS(items...).(*array)),
		less,
	}
}

func (a *arraySortTS) Less(i, j int) bool {
	a.arrayTS.l.RLock()
	defer a.arrayTS.l.RUnlock()
	return a.less(i, j)
}

func (a *arraySortTS) Sort() {
	a.arrayTS.l.RLock()
	defer a.arrayTS.l.RUnlock()
	sort.Sort(sort.Interface(a))
}
