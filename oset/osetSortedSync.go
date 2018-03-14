package oset

import (
	"sort"
)

// Sorted is the interface for sortable osets
type osetSortSync struct {
	*osetSync
	less func(slice []interface{}, i, j int) bool
}

// NewSortedSync creates an ordered  thread safe set that expose Sort method
func NewSortedSync(less func(slice []interface{}, i, j int) bool, items ...interface{}) Sorted {
	return &osetSortSync{
		(NewSync(items...).(*osetSync)),
		less,
	}
}

func (a *osetSortSync) Less(i, j int) bool {
	a.osetSync.l.RLock()
	defer a.osetSync.l.RUnlock()
	return a.less(a.Slice(), i, j)
}

func (a *osetSortSync) Sort() {
	sort.Sort(a)
}
