package util

type StringSet map[string]interface{}

func (s *StringSet) Add(str ...string) {
	for _, o := range str {
		(*s)[o] = str
	}
}

func (s *StringSet) AddAll(other StringSet) {
	for o := range other {
		s.Add(o)
	}
}

func (s *StringSet) RemoveAll(other StringSet) {
	for o := range other {
		s.Remove(o)
	}
}

func (s *StringSet) Remove(str ...string) int {
	count := 0
	for _, o := range str {
		if _, ok := (*s)[o]; ok {
			count++
			delete(*s, o)
		}
	}
	return count
}

func (s *StringSet) Has(str string) bool {
	_, ok := (*s)[str]
	return ok
}

func (s *StringSet) Retain(str ...string) int {
	o := NewStringSet(str...)

	count := 0
	for k := range *s {
		if !o.Has(k) {
			count++
			delete(*s, k)
		}
	}
	return count
}

func (s *StringSet) Slice() []string {
	l := make([]string, 0, len(*s))
	for s2 := range *s {
		l = append(l, s2)
	}
	return l
}

func NewStringSet(str ...string) StringSet {
	s := StringSet{}
	for _, o := range str {
		s[o] = o
	}
	return s
}
