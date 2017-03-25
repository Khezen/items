package set

import (
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestSetTSNew(t *testing.T) {
	s := NewTS()

	if s.Len() != 0 {
		t.Error("New: calling without any parameters should create a set with zero size")
	}

}

func TestSetTSNewParameters(t *testing.T) {
	s := NewTS()
	s.Add("string", "another_string", 1, 3.14)

	if s.Len() != 4 {
		t.Error("New: calling with parameters should create a set with size of four")
	}
}

func TestSetTSAdd(t *testing.T) {
	s := NewTS()
	s.Add(1)
	s.Add(2)
	s.Add(2) // duplicate
	s.Add("fatih")
	s.Add("zeynep")
	s.Add("zeynep") // another duplicate

	if s.Len() != 4 {
		t.Error("Add: items are not unique. The set size should be four")
	}

	if !s.Has(1, 2, "fatih", "zeynep") {
		t.Error("Add: added items are not availabile in the set.")
	}
}

func TestSetTSAddMultiple(t *testing.T) {
	s := NewTS()
	s.Add("ankara", "san francisco", 3.14)

	if s.Len() != 3 {
		t.Error("Add: items are not unique. The set size should be three")
	}

	if !s.Has("ankara", "san francisco", 3.14) {
		t.Error("Add: added items are not availabile in the set.")
	}
}

func TestSetTSRemove(t *testing.T) {
	s := NewTS()
	s.Add(1)
	s.Add(2)
	s.Add("fatih")

	s.Remove(1)
	if s.Len() != 2 {
		t.Error("Remove: set size should be two after removing")
	}

	s.Remove(1)
	if s.Len() != 2 {
		t.Error("Remove: set size should be not change after trying to remove a non-existing item")
	}

	s.Remove(2)
	s.Remove("fatih")
	if s.Len() != 0 {
		t.Error("Remove: set size should be zero")
	}

	s.Remove("fatih") // try to remove something from a zero length set
}

func TestSetTSRemoveMultiple(t *testing.T) {
	s := NewTS()
	s.Add("ankara", "san francisco", 3.14, "istanbul")
	s.Remove("ankara", "san francisco", 3.14)

	if s.Len() != 1 {
		t.Error("Remove: items are not unique. The set size should be four")
	}

	if !s.Has("istanbul") {
		t.Error("Add: added items are not availabile in the set.")
	}
}

func TestSetTS_Pop(t *testing.T) {
	s := NewTS()
	s.Add(1)
	s.Add(2)
	s.Add("fatih")

	a := s.Pop()
	if s.Len() != 2 {
		t.Error("Pop: set size should be two after popping out")
	}

	if s.Has(a) {
		t.Error("Pop: returned item should not exist")
	}

	s.Pop()
	s.Pop()
	b := s.Pop()
	if b != nil {
		t.Error("Pop: should return nil because set is empty")
	}

	s.Pop() // try to remove something from a zero length set
}

func TestSetTSHas(t *testing.T) {
	s := NewTS()
	s.Add("1", "2", "3", "4")

	if !s.Has("1") {
		t.Error("Has: the item 1 exist, but 'Has' is returning false")
	}

	if !s.Has("1", "2", "3", "4") {
		t.Error("Has: the items all exist, but 'Has' is returning false")
	}
}

func TestSetTSClear(t *testing.T) {
	s := NewTS()
	s.Add(1)
	s.Add("istanbul")
	s.Add("san francisco")

	s.Clear()
	if s.Len() != 0 {
		t.Error("Clear: set size should be zero")
	}
}

func TestSetTS_IsEmpty(t *testing.T) {
	s := NewTS()

	empty := s.IsEmpty()
	if !empty {
		t.Error("IsEmpty: set is empty, it should be true")
	}

	s.Add(2)
	s.Add(3)
	notEmpty := s.IsEmpty()

	if notEmpty {
		t.Error("IsEmpty: set is filled, it should be false")
	}
}

func TestSetTSIsEqual(t *testing.T) {
	// same size, same content
	s := NewTS()
	s.Add("1", "2", "3")
	u := NewTS()
	u.Add("1", "2", "3")

	ok := s.IsEqual(u)
	if !ok {
		t.Error("IsEqual: set s and t are equal. However it returns false")
	}

	// same size, different content
	a := NewTS()
	a.Add("1", "2", "3")
	b := NewTS()
	b.Add("4", "5", "6")

	ok = a.IsEqual(b)
	if ok {
		t.Error("IsEqual: set a and b are now equal (1). However it returns true")
	}

	// different size, similar content
	a = NewTS()
	a.Add("1", "2", "3")
	b = NewTS()
	b.Add("1", "2", "3", "4")

	ok = a.IsEqual(b)
	if ok {
		t.Error("IsEqual: set s and t are now equal (2). However it returns true")
	}
}

func TestSetTSIsSubset(t *testing.T) {
	s := NewTS()
	s.Add("1", "2", "3", "4")
	u := NewTS()
	u.Add("1", "2", "3")

	ok := s.IsSubset(u)
	if !ok {
		t.Error("IsSubset: u is a subset of s. However it returns false")
	}

	ok = u.IsSubset(s)
	if ok {
		t.Error("IsSubset: s is not a subset of u. However it returns true")
	}

}

func TestSetTSIsSuperset(t *testing.T) {
	s := NewTS()
	s.Add("1", "2", "3", "4")
	u := NewTS()
	u.Add("1", "2", "3")

	ok := u.IsSuperset(s)
	if !ok {
		t.Error("IsSuperset: s is a superset of u. However it returns false")
	}

	ok = s.IsSuperset(u)
	if ok {
		t.Error("IsSuperset: u is not a superset of u. However it returns true")
	}

}

func TestSetTSString(t *testing.T) {
	s := NewTS()
	if s.String() != "[]" {
		t.Errorf("String: output is not what is excepted '%s'", s.String())
	}

	if !strings.HasPrefix(s.String(), "[") {
		t.Error("String: output should begin with a square bracket")
	}

	if !strings.HasSuffix(s.String(), "]") {
		t.Error("String: output should end with a square bracket")
	}
}

func TestSetTS_Slice(t *testing.T) {
	s := NewTS()
	s.Add("1", "2", "3", "4")

	// this returns a slice of interface{}
	if len(s.Slice()) != 4 {
		t.Error("Slice: slice size should be four.")
	}

	for _, item := range s.Slice() {
		r := reflect.TypeOf(item)
		if r.Kind().String() != "string" {
			t.Error("Slice: slice item should be a string")
		}
	}
}

func TestSetTSCopy(t *testing.T) {
	s := NewTS()
	s.Add("1", "2", "3", "4")
	r := s.Copy()

	if !s.IsEqual(r) {
		t.Error("Copy: set s and r are not equal")
	}
}

func TestSetTSMerge(t *testing.T) {
	s := NewTS()
	s.Add("1", "2", "3")
	r := NewTS()
	r.Add("3", "4", "5")
	s.Merge(r)

	if s.Len() != 5 {
		t.Error("Merge: the set doesn't have all items in it.")
	}

	if !s.Has("1", "2", "3", "4", "5") {
		t.Error("Merge: merged items are not availabile in the set.")
	}
}

func TestSetTSSeparate(t *testing.T) {
	s := NewTS()
	s.Add("1", "2", "3")
	r := NewTS()
	r.Add("3", "5")
	s.Separate(r)

	if s.Len() != 2 {
		t.Error("Separate: the set doesn't have all items in it.")
	}

	if !s.Has("1", "2") {
		t.Error("Separate: items after separation are not availabile in the set.")
	}
}

func TestSetTS_RaceAdd(t *testing.T) {
	// Create two sets. Add concurrently items to each of them. Remove from the
	// other one.
	// "go test -race" should detect this if the library is not thread-safe.
	s := NewTS()
	u := NewTS()

	go func() {
		for i := 0; i < 1000; i++ {
			item := "item" + strconv.Itoa(i)
			go func(i int) {
				s.Add(item)
				u.Add(item)
			}(i)
		}
	}()

	for i := 0; i < 1000; i++ {
		item := "item" + strconv.Itoa(i)
		go func(i int) {
			s.Add(item)
			u.Add(item)
		}(i)
	}
}
