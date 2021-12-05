package set

type Ints map[int]interface{}

func (s *Ints) Add(i ...int) {
	for _, o := range i {
		(*s)[o] = i
	}
}

func (s *Ints) AddAll(other Ints) {
	for o := range other {
		s.Add(o)
	}
}

func (s *Ints) RemoveAll(other Ints) {
	for o := range other {
		s.Remove(o)
	}
}

func (s *Ints) Remove(i ...int) {
	for _, o := range i {
		delete(*s, o)
	}
}

func NewInts(i ...int) Ints {
	s := Ints{}
	for _, o := range i {
		s[o] = o
	}
	return s
}

func (s *Ints) Contains(i int) bool {
	_, ok := (*s)[i]
	return ok
}

func (s *Ints) ContainsSlice(i []int) bool {
	for _, v := range i {
		if !s.Contains(v) {
			return false
		}
	}
	return true
}
