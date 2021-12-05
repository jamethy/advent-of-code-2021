package set

type Strings map[string]interface{}

func (s *Strings) Add(str ...string) {
	for _, o := range str {
		(*s)[o] = str
	}
}

func (s *Strings) AddAll(other Strings) {
	for o := range other {
		s.Add(o)
	}
}

func (s *Strings) RemoveAll(other Strings) {
	for o := range other {
		s.Remove(o)
	}
}

func (s *Strings) Remove(str ...string) int {
	count := 0
	for _, o := range str {
		if _, ok := (*s)[o]; ok {
			count++
			delete(*s, o)
		}
	}
	return count
}

func (s *Strings) Has(str string) bool {
	_, ok := (*s)[str]
	return ok
}

func (s *Strings) Retain(str ...string) int {
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

func (s *Strings) Slice() []string {
	l := make([]string, 0, len(*s))
	for s2 := range *s {
		l = append(l, s2)
	}
	return l
}

func NewStringSet(str ...string) Strings {
	s := Strings{}
	for _, o := range str {
		s[o] = o
	}
	return s
}
