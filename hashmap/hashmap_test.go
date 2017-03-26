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

}

func TestHas(t *testing.T) {

}

func TestHasValue(t *testing.T) {

}

func TestKeyOf(t *testing.T) {

}

func TestEach(t *testing.T) {

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

func TestCopy(t *testing.T) {

}
