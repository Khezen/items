package array

import (
	"sort"
)

// Sorted is the interface for sortable arrays
type arraySortSync struct {
	*arraySync
	less func(slice []interface{}, i, j int) bool
}

// NewSortedSync creates a thread safe array that expose Sort method
func NewSortedSync(less func(slice []interface{}, i, j int) bool, items ...interface{}) Sorted {
	return &arraySortSync{
		(NewSync(items...).(*arraySync)),
		less,
	}
}

func (a *arraySortSync) Less(i, j int) bool {
	a.arraySync.l.RLock()
	defer a.arraySync.l.RUnlock()
	return a.less(a.arraySync.s, i, j)
}

func (a *arraySortSync) Sort() {
	sort.Sort(a)
}
