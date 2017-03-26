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
		h                 Interface
		counter, expected int
	}{
		{New("1", 1, "-8", -8, "42", 42), 0, 2},
		{NewTS("1", 1, "-8", -8, "42", 42), 0, 2},
	}
	for _, c := range cases {
		callback := func(k, v interface{}) bool {
			value := v.(int)
			c.counter++
			return value > 0
		}
		c.h.Each(callback)
		if c.counter != c.expected {
			t.Errorf("Expected %v. Got %v.", c.expected, c.counter)
		}
	}
}

func TestLen(t *testing.T) {

}

func TestClear(t *testing.T) {

}

func TestIsEmpty(t *testing.T) {

}

func TestIsEqual(t *testing.T) {

}

func TestSting(t *testing.T) {

}

func TestKeys(t *testing.T) {

}

func TestValues(t *testing.T) {

}

func TestMap(t *testing.T) {

}
func TestCopy(t *testing.T) {

}
