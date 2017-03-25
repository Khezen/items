package set

import (
	"reflect"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	s := New()
	s.Add(1, 2, 3, "testing")
	if s.Len() != 4 {
		t.Error("New: The set created was expected have 4 items")
	}
}

func TestSetAdd(t *testing.T) {
	s := New()
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

func TestSetAddMultiple(t *testing.T) {
	s := New()
	s.Add("ankara", "san francisco", 3.14)

	if s.Len() != 3 {
		t.Error("Add: items are not unique. The set size should be three")
	}

	if !s.Has("ankara", "san francisco", 3.14) {
		t.Error("Add: added items are not availabile in the set.")
	}
}

func TestSetRemove(t *testing.T) {
	s := New()
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

func TestSetRemoveMultiple(t *testing.T) {
	s := New()
	s.Add("ankara", "san francisco", 3.14, "istanbul")
	s.Remove("ankara", "san francisco", 3.14)

	if s.Len() != 1 {
		t.Error("Remove: items are not unique. The set size should be four")
	}

	if !s.Has("istanbul") {
		t.Error("Add: added items are not availabile in the set.")
	}
}

func TestSetPop(t *testing.T) {
	s := New()
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

func TestSetHas(t *testing.T) {
	s := New()
	s.Add("1", "2", "3", "4")

	if !s.Has("1") {
		t.Error("Has: the item 1 exist, but 'Has' is returning false")
	}

	if !s.Has("1", "2", "3", "4") {
		t.Error("Has: the items all exist, but 'Has' is returning false")
	}
}

func TestSetClear(t *testing.T) {
	s := New()
	s.Add(1)
	s.Add("istanbul")
	s.Add("san francisco")

	s.Clear()
	if s.Len() != 0 {
		t.Error("Clear: set size should be zero")
	}
}

func TestSetIsEmpty(t *testing.T) {
	s := New()
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

func TestSetIsEqual(t *testing.T) {
	s, u, v := New("1", "2", "3"), New("1", "2", "3"), NewTS("1", "2", "3")
	ok := s.IsEqual(u)
	if !ok {
		t.Error("IsEqual: set s and t are equal. However it returns false")
	}
	ok = s.IsEqual(v)
	if !ok {
		t.Error("IsEqual: set s and t are equal. However it returns false")
	}

	// same size, different content

	a, b, c := New("1", "2", "3"), New("4", "5", "6"), NewTS("4", "5", "6")
	ok = a.IsEqual(b)
	if ok {
		t.Error("IsEqual: set a and b are now equal (1). However it returns true")
	}
	ok = a.IsEqual(c)
	if ok {
		t.Error("IsEqual: set a and b are now equal (1). However it returns true")
	}

	// different size, similar content
	a, b, c = New("1", "2", "3"), New("1", "2", "3", "4"), NewTS("1", "2", "3", "4")

	ok = a.IsEqual(b)
	if ok {
		t.Error("IsEqual: set s and t are now equal (2). However it returns true")
	}
	ok = a.IsEqual(c)
	if ok {
		t.Error("IsEqual: set s and t are now equal (2). However it returns true")
	}
}

func TestSetIsSubset(t *testing.T) {
	s := New()
	s.Add("1", "2", "3", "4")
	u := New()
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

func TestSetIsSuperset(t *testing.T) {
	s := New()
	s.Add("1", "2", "3", "4")
	u := New()
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

func TestSetString(t *testing.T) {
	s := New()
	if s.String() != "[]" {
		t.Errorf("String: output is not what is excepted '%s'", s.String())
	}

	s.Add("1", "2", "3", "4")

	if !strings.HasPrefix(s.String(), "[") {
		t.Error("String: output should begin with a square bracket")
	}

	if !strings.HasSuffix(s.String(), "]") {
		t.Error("String: output should end with a square bracket")
	}
}

func TestSetSlice(t *testing.T) {
	s := New()
	s.Add("1", "2", "3", "4")
	s = New()
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

func TestSetCopy(t *testing.T) {
	s := New()
	s.Add("1", "2", "3", "4")
	r := s.Copy()

	if !s.IsEqual(r) {
		t.Error("Copy: set s and r are not equal")
	}
}

func TestSetMerge(t *testing.T) {
	s := New()
	s.Add("1", "2", "3")
	r := New()
	r.Add("3", "4", "5")
	s.Merge(r)

	if s.Len() != 5 {
		t.Error("Merge: the set doesn't have all items in it.")
	}

	if !s.Has("1", "2", "3", "4", "5") {
		t.Error("Merge: merged items are not availabile in the set.")
	}
}

func TestSetSeparate(t *testing.T) {
	s := New()
	s.Add("1", "2", "3")
	r := New()
	r.Add("3", "5")
	s.Separate(r)

	if s.Len() != 2 {
		t.Error("Separate: the set doesn't have all items in it.")
	}

	if !s.Has("1", "2") {
		t.Error("Separate: items after separation are not availabile in the set.")
	}
}

func TestRetain(t *testing.T) {
	cases := []struct {
		set, toBeRetained, expected Interface
	}{
		{New(1, 2, 3, 4), New(2, 4, 8), New(2, 4)},
		{New(1, 2, 3, 4), New(2, 8), New(2)},
		{New(1, 2, 3, 4), NewTS(2, 4, 8), New(2, 4)},
		{New(1, 2, 3, 4), NewTS(2, 8), New(2)},
		{NewTS(1, 2, 3, 4), New(2, 4, 8), NewTS(2, 4)},
		{NewTS(1, 2, 3, 4), New(2, 8), NewTS(2)},
		{NewTS(1, 2, 3, 4), NewTS(2, 4, 8), NewTS(2, 4)},
		{NewTS(1, 2, 3, 4), NewTS(2, 8), NewTS(2)},
	}
	for _, c := range cases {
		c.set.Retain(c.toBeRetained)
		if !c.set.IsEqual(c.expected) {
			t.Errorf("Expected %v. Got %v", c.expected.Slice(), c.set.Slice())
		}
	}
}
