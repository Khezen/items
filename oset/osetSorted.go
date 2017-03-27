package oset

// Sortable is the interface for sortable ordered sets
type Sortable interface {
	Interface
	Sort()
	Less(i, j int) bool
}
