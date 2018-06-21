package goarray

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

func (g *Goarray) Filter(c func(a int) bool) (result Goarray) {
	for _, elem := range *g {
		if c(elem) {
			result = append(result, elem)
		}
	}
	return
}

func (g *Goarray) Map(funcao func(a int) int) (result Goarray) {
	for _, elem := range *g {
		result = append(result, funcao(elem))
	}
	return
}
