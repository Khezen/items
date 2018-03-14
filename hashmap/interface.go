package hashmap

// Interface describes functions a Map must expose
type Interface interface {
	Get(k interface{}) (interface{}, error)
	Put(k, v interface{})
	Remove(keys ...interface{})
	Has(keys ...interface{}) bool
	HasValue(values ...interface{}) bool
	KeyOf(value interface{}) (interface{}, error)
	Each(func(k, v interface{}) bool)

	Len() int
	Clear()
	IsEmpty() bool
	IsEqual(Interface) bool

	String() string
	Keys() []interface{}
	Values() []interface{}
	Map() map[interface{}]interface{}
	Copy() Interface
}
