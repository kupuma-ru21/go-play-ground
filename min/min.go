package min

import "sort"

type Miner interface {
	Len() int
	ElemIx(ix int) interface{}
	Less(i, j int) bool
}

func Min(data Miner) interface{} {
	switch v := data.(type) {
	case IntArray:
		sort.Ints(v)
		return v[0]
	case StringArray:
		sort.Strings(v)
		return v[0]
	default:
		return "unknown type"
	}
}

type IntArray []int

func (p IntArray) Len() int {
	return len(p)
}
func (p IntArray) ElemIx(ix int) interface{} {
	return p[ix]
}
func (p IntArray) Less(i, j int) bool {
	return p[i] > p[j]
}

type StringArray []string

func (p StringArray) Len() int {
	return len(p)
}
func (p StringArray) ElemIx(ix int) interface{} {
	return p[ix]
}
func (p StringArray) Less(i, j int) bool {
	return p[i] > p[j]
}
