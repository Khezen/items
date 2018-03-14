package oset

// Sorted is the interface for sortable ordered sets
type Sorted interface {
	Interface
	Sort()
	Less(i, j int) bool
}
