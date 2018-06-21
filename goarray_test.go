package goarray

import (
	"testing"
)

func TestEquals(t *testing.T) {
	a := Goarray{1, 2, 3}
	b := Goarray{1, 2, 3}

	if !a.Equals(b) {
		t.Errorf("First array: %v \n Second array: %v \n The result cound't be false", a, b)
	}

	a = Goarray{}
	b = Goarray{}

	if !b.Equals(a) {
		t.Errorf("First array: %v \n Second array: %v \n The result cound't be false", a, b)
	}

	a = Goarray{1, 2, 3, 4, 5}
	b = Goarray{0, 2, 3}

	if b.Equals(a) {
		t.Errorf("First array: %v \n Second array: %v \n The result cound't be true", a, b)
	}

	a = Goarray{1, 2, 3, 4, 5}
	b = Goarray{1, 2, 3, 5, 4}

	if b.Equals(a) {
		t.Errorf("First array: %v \n Second array: %v \n The result cound't be true", a, b)
	}
}

func TestFilter(t *testing.T) {
	b := Goarray{1, 2, 3, 4, 5, 6, 7}

	result := b.Filter(func(elem int) bool {
		if elem%2 > 0 {
			return true
		}
		return false
	})

	expected := Goarray{1, 3, 5, 7}
	if !result.Equals(expected) {
		t.Errorf("Expected value: %v \n Returned value: %v", expected, result)
	}
}

func TestMap(t *testing.T) {
	b := Goarray{1, 2, 3, 4, 5, 6, 7}

	result := b.Map(func(elem int) int {
		return elem * 2
	})

	expected := Goarray{2, 4, 6, 8, 10, 12, 14}
	if !result.Equals(expected) {
		t.Errorf("Expected value: %v \n Returned value: %v", expected, result)
	}
}

func TestJoin(t *testing.T) {
	b := Goarray{1, 2, 3, 4, 5}

	result := b.Join(",")

	expected := "1,2,3,4,5"
	if result != expected {
		t.Errorf("Expected value: %v \n Returned value: %v", expected, result)
	}
}

func TestIncludes(t *testing.T) {
	b := Goarray{1, 2, 3, 4, 5}

	result := b.Includes(2)

	if !result {
		t.Errorf("Expected value: %t \n Returned value: %t", true, result)
	}

	result = b.Includes(20)

	if result {
		t.Errorf("Expected value: %t \n Returned value: %t", false, result)
	}
}

func TestIndexOf(t *testing.T) {
	b := Goarray{1, 2, 3, 4, 5}

	result := b.IndexOf(2)
	expect := 1
	if expect != result {
		t.Errorf("Expected value: %d \n Returned value: %d", expect, result)
	}

	result = b.IndexOf(20)
	expect = -1
	if expect != result {
		t.Errorf("Expected value: %d \n Returned value: %d", expect, result)
	}
}
