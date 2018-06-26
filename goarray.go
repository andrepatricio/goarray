package goarray

import (
	"fmt"
)

type Goarray []int

func (g *Goarray) Equals(b Goarray) bool {
	if len(*g) != len(b) {
		return false
	}
	for n, elem := range *g {
		if elem != b[n] {
			return false
		}
	}
	return true
}

// Creates a new array with all elements that pass the test implemented by the provided function.
func (g *Goarray) Filter(c func(a int) bool) (result Goarray) {
	for _, elem := range *g {
		if c(elem) {
			result = append(result, elem)
		}
	}
	return
}

// Creates a new array with the results of calling a provided function on every element in the calling array.
func (g *Goarray) Map(funcao func(a int) int) (result Goarray) {
	for _, elem := range *g {
		result = append(result, funcao(elem))
	}
	return
}

// Joins all elements of an array into a string and returns this string.
func (g *Goarray) Join(s string) (result string) {
	for n, elem := range *g {
		switch {
		case n == len(*g)-1:
			result += fmt.Sprintf("%d", elem)
		default:
			result += fmt.Sprintf("%d%s", elem, s)
		}
	}
	return
}

// Determines whether an array includes a certain element,
// returning true or false as appropriate.
func (g *Goarray) Includes(n int) bool {
	for _, elem := range *g {
		if elem == n {
			return true
		}
	}
	return false
}

// Returns the first index at which a given element can be found in the array,
// or -1 if it is not present.
func (g *Goarray) IndexOf(n int) int {
	for i, elem := range *g {
		if n == elem {
			return i
		}
	}
	return -1
}

// Returns the last index at which a given element can be found in the array,
// or -1 if it is not present.
func (g *Goarray) LastIndexOf(n int) int {
	result := -1
	for i, elem := range *g {
		if n == elem {
			result = i
		}
	}
	return result
}

// Tests whether all elements in the array pass
// the test implemented by the provided function.
func (g *Goarray) Every(test func(n int) bool) bool {
	for _, elem := range *g {
		if !test(elem) {
			return false
		}
	}
	return true
}

// Method reverses an array. Return a new array without change the first one
func (g *Goarray) Reverse() Goarray {
	gr := make(Goarray, len(*g))
	for n, elem := range *g {
		gr[(len(*g)-1)-n] = elem
	}
	return gr
}
