package array

// Sortable is the interface for sortable arrays
type Sortable interface {
	Interface
	Sort()
	Less(i, j int) bool
}
