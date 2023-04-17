package main

// реализуйте быстрое множество
type IntSet struct {
	set map[int]bool
}

func MakeIntSet() IntSet {
	return IntSet{set: make(map[int]bool)}
}

func (s IntSet) Contains(elem int) bool {
	return s.set[elem]
}

func (s *IntSet) Add(elem int) bool {
	if s.Contains(elem) {
		return false
	}
	s.set[elem] = true
	return true
}
