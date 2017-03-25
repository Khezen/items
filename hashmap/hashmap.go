package hashmap

// Interface describes functions a Map must expose
type Interface interface {
	Get(k interface{})
	Put(k, v interface{})
	Remove(keys ...interface{})
	Has(keys ...interface{}) bool
	HasValue(values ...interface{}) bool
	Each(func(k, v interface{}) bool)

	Len() int
	Clear()
	IsEmpty() bool
	IsEqual(Interface) bool

	String() string
	Keys() []interface{}
	Values() []interface{}
	Copy() Interface
}
