package hashmap

import (
	"github.com/khezen/check"
	"testing"
)

func testErr(err error, expectErr bool, t *testing.T) {
	if !check.ErrorExpectation(err, expectErr) {
		t.Errorf(" Error expected? %v. Got: %v.", expectErr, err)
	}
}

func TestGet(t *testing.T) {
	cases := []struct {
		h             Interface
		key, expected interface{}
		expectErr     bool
	}{
		{New("1", 1, "42", 42, "-8", -8), "-8", -8, false},
		{New("1", 1, "42", 42, "-8", -8), "1000", nil, true},
		{NewTS("1", 1, "42", 42, "-8", -8), "-8", -8, false},
		{NewTS("1", 1, "42", 42, "-8", -8), "1000", nil, true},
	}
	for _, c := range cases {
		item, err := c.h.Get(c.key)
		testErr(err, c.expectErr, t)
		if item != c.expected {
			t.Errorf("Expected %v. Got %v.", c.expected, item)
		}
	}
}

func TestPut(t *testing.T) {
	cases := []struct {
		h, expected Interface
		key, value  interface{}
	}{
		{New("1", 1, "42", 42), New("1", 1, "42", 42, "-8", -8), "-8", -8},
	}
	for _, c := range cases {
		c.h.Put(c.key, c.value)
		if !c.h.IsEqual(c.expected) {
			t.Errorf("Expected %v. Got %v.", c.expected.Map(), c.h.Map())
		}
	}
}

func TestRemove(t *testing.T) {
	cases := []struct {
		h, expected Interface
		keys        []interface{}
	}{
		{New("1", 1, "42", 42, "-8", -8), New("1", 1, "42", 42), []interface{}{"-8"}},
		{New("1", 1, "42", 42, "-8", -8), New("42", 42), []interface{}{"-8", "1"}},
		{New("1", 1, "42", 42, "-8", -8), New("1", 1, "42", 42), []interface{}{"-8", "-1"}},
		{NewTS("1", 1, "42", 42, "-8", -8), NewTS("1", 1, "42", 42), []interface{}{"-8"}},
		{NewTS("1", 1, "42", 42, "-8", -8), NewTS("42", 42), []interface{}{"-8", "1"}},
		{NewTS("1", 1, "42", 42, "-8", -8), NewTS("1", 1, "42", 42), []interface{}{"-8", "-1"}},
	}
	for _, c := range cases {
		c.h.Remove(c.keys...)
		if !c.h.IsEqual(c.expected) {
			t.Errorf("Expected %v. Got %v.", c.expected.Map(), c.h.Map())
		}
	}
}

func TestHas(t *testing.T) {
	cases := []struct {
		h    Interface
		keys []interface{}
		has  bool
	}{
		{New("1", 1, "42", 42, "-8", -8), []interface{}{"-8"}, true},
		{New("1", 1, "42", 42, "-8", -8), []interface{}{"-8", "1"}, true},
		{New("1", 1, "42", 42, "-8", -8), []interface{}{"-8", "-1"}, false},
		{NewTS("1", 1, "42", 42, "-8", -8), []interface{}{"-8"}, true},
		{NewTS("1", 1, "42", 42, "-8", -8), []interface{}{"-8", "1"}, true},
		{NewTS("1", 1, "42", 42, "-8", -8), []interface{}{"-8", "-1"}, false},
	}
	for _, c := range cases {
		if c.h.Has(c.keys...) != c.has {
			t.Errorf("Expected %v. Got %v. => %v", c.has, c.h.Has(c.keys...), c.h.String())
		}
	}
}

func TestHasValue(t *testing.T) {
	cases := []struct {
		h      Interface
		values []interface{}
		has    bool
	}{
		{New("1", 1, "42", 42, "-8", -8), []interface{}{-8}, true},
		{New("1", 1, "42", 42, "-8", -8), []interface{}{-8, 1}, true},
		{New("1", 1, "42", 42, "-8", -8), []interface{}{-8, -1}, false},
		{NewTS("1", 1, "42", 42, "-8", -8), []interface{}{-8}, true},
		{NewTS("1", 1, "42", 42, "-8", -8), []interface{}{-8, 1}, true},
		{NewTS("1", 1, "42", 42, "-8", -8), []interface{}{-8, -1}, false},
	}
	for _, c := range cases {
		if c.h.HasValue(c.values...) != c.has {
			t.Errorf("Expected %v. Got %v. => %v", c.has, c.h.HasValue(c.values...), c.h.String())
		}
	}
}

func TestKeyOf(t *testing.T) {
	cases := []struct {
		h          Interface
		value, key interface{}
		expectErr  bool
	}{
		{New("1", 1, "42", 42, "-8", -8), -8, "-8", false},
		{New("1", 1, "42", 42, "-8", -8), 1000, nil, true},
		{NewTS("1", 1, "42", 42, "-8", -8), -8, "-8", false},
		{NewTS("1", 1, "42", 42, "-8", -8), 1000, nil, true},
	}
	for _, c := range cases {
		key, err := c.h.KeyOf(c.value)
		testErr(err, c.expectErr, t)
		if key != c.key {
			t.Errorf("Expected %v. Got %v. => %v", c.key, key, c.h.String())
		}
	}
}

func TestEach(t *testing.T) {
	cases := []struct {
		h    Interface
		stop bool
	}{
		{New("1", 1, "-8", -8, "42", 42), true},
		{NewTS("1", 1, "-8", -8, "42", 42), false},
	}
	for _, c := range cases {
		count := 0
		callback := func(k, v interface{}) bool {
			count++
			return !c.stop
		}
		c.h.Each(callback)
		if c.stop && count != 1 {
			t.Errorf("Expected %v. Got %v.", 1, count)
		}
	}
}

func TestLen(t *testing.T) {
	cases := []struct {
		h   Interface
		len int
	}{
		{New("1", 1, "-8", -8, "42", 42), 3},
		{NewTS("1", 1, "-8", -8, "42", 42), 3},
	}
	for _, c := range cases {
		length := c.h.Len()
		if length != c.len {
			t.Errorf("Expected %v. Got %v.", c.len, length)
		}
	}
}

func TestClear(t *testing.T) {
	cases := []struct {
		h Interface
	}{
		{New("1", 1, "-8", -8, "42", 42)},
		{NewTS("1", 1, "-8", -8, "42", 42)},
	}
	for _, c := range cases {
		c.h.Clear()
		if !c.h.IsEmpty() {
			t.Errorf("map %v should be empty", c.h.String())
		}
	}
}

func TestIsEmpty(t *testing.T) {
	cases := []struct {
		h     Interface
		empty bool
	}{
		{New("1", 1, "-8", -8, "42", 42), false},
		{New(), true},
		{NewTS("1", 1, "-8", -8, "42", 42), false},
		{NewTS(), true},
	}
	for _, c := range cases {
		empty := c.h.IsEmpty()
		if empty != c.empty {
			t.Errorf("Expected %v. Got %v.", c.empty, empty)
		}
	}
}

func TestIsEqual(t *testing.T) {
	cases := []struct {
		h, t  Interface
		equal bool
	}{
		{New("1", 1, "-8", -8, "42", 42), New("1", 1, "-8", -8), false},
		{New("1", 1, "-8", -8, "42", 42), New("1", 1, "-8", -8, "42"), false},
		{New("1", 1, "-8", -8, "42", 42), New("1", 1, "-8", -8, "42", 42), true},
		{NewTS("1", 1, "-8", -8, "42", 42), NewTS("1", 1, "-8", -8), false},
		{NewTS("1", 1, "-8", -8, "42", 42), NewTS("1", 1, "-8", -8, "42"), false},
		{NewTS("1", 1, "-8", -8, "42", 42), NewTS("1", 1, "-8", -8, "42", 42), true},
	}
	for _, c := range cases {
		equal := c.h.IsEqual(c.t)
		if equal != c.equal {
			t.Errorf("Expected %v. Got %v.", c.equal, equal)
		}
	}
}

func TestString(t *testing.T) {
	cases := []struct {
		h   Interface
		str string
	}{
		{New("1", 1, "-8", -8, "42", 42), "[1:1 -8:-8 42:42]"},
		{NewTS("1", 1, "-8", -8, "42", 42), "[1:1 -8:-8 42:42]"},
	}
	for _, c := range cases {
		str := c.h.String()
		if str != c.str {
			t.Errorf("Expected %v. Got %v.", c.str, str)
		}
	}
}

func TestKeys(t *testing.T) {

}

func TestValues(t *testing.T) {

}

func TestMap(t *testing.T) {

}
func TestCopy(t *testing.T) {

}
