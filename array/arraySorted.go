package array

// Sorted is the interface for sortable arrays
type Sorted interface {
	Interface
	Sort()
	Less(i, j int) bool
}
