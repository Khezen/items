package oset

import (
	"sort"
)

// Sorted is the interface for sortable osets
type osetSortTS struct {
	*osetTS
	less func(slice []interface{}, i, j int) bool
}

// NewSortedTS creates an ordered  thread safe set that expose Sort method
func NewSortedTS(less func(slice []interface{}, i, j int) bool, items ...interface{}) Sorted {
	return &osetSortTS{
		(NewTS(items...).(*osetTS)),
		less,
	}
}

func (a *osetSortTS) Less(i, j int) bool {
	a.osetTS.l.RLock()
	defer a.osetTS.l.RUnlock()
	return a.less(a.Slice(), i, j)
}

func (a *osetSortTS) Sort() {
	sort.Sort(a)
}
