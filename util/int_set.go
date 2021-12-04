package util

type IntSet map[int]interface{}

func (s *IntSet) Add(i ...int) {
	for _, o := range i {
		(*s)[o] = i
	}
}

func (s *IntSet) AddAll(other IntSet) {
	for o := range other {
		s.Add(o)
	}
}

func (s *IntSet) RemoveAll(other IntSet) {
	for o := range other {
		s.Remove(o)
	}
}

func (s *IntSet) Remove(i ...int) {
	for _, o := range i {
		delete(*s, o)
	}
}

func NewIntSet(i ...int) IntSet {
	s := IntSet{}
	for _, o := range i {
		s[o] = o
	}
	return s
}

func (s *IntSet) Contains(i int) bool {
	_, ok := (*s)[i]
	return ok
}

func (s *IntSet) ContainsSlice(i []int) bool {
	for _, v := range i {
		if !s.Contains(v) {
			return false
		}
	}
	return true
}
