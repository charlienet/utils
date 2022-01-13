package set

import "github.com/charlienet/utils/lang"

type threadUnsafeSet map[interface{}]lang.PlaceholderType

func newThreadUnsafeSet() threadUnsafeSet {
	return make(threadUnsafeSet)
}

func (s *threadUnsafeSet) Add(e interface{}) bool {
	if _, found := (*s)[e]; found {
		return false
	}

	(*s)[e] = lang.Placeholder
	return true
}

func (s *threadUnsafeSet) Clear() {
	*s = newThreadUnsafeSet()
}

func (s *threadUnsafeSet) Remove(i interface{}) {
	delete(*s, i)
}

func (s *threadUnsafeSet) Clone() Set {
	clonedSet := newThreadUnsafeSet()
	for elem := range *s {
		clonedSet.Add(elem)
	}

	return &clonedSet
}

func (s *threadUnsafeSet) Contains(i ...interface{}) bool {
	for _, val := range i {
		if _, ok := (*s)[val]; !ok {
			return false
		}
	}
	return true
}

func (s *threadUnsafeSet) Union(other Set) Set {
	o := other.(*threadUnsafeSet)

	unionedSet := newThreadUnsafeSet()

	for elem := range *s {
		unionedSet.Add(elem)
	}
	for elem := range *o {
		unionedSet.Add(elem)
	}
	return &unionedSet
}

func (s *threadUnsafeSet) Intersect(other Set) Set {
	o := other.(*threadUnsafeSet)

	intersection := newThreadUnsafeSet()
	// loop over smaller set
	if s.Count() < other.Count() {
		for elem := range *s {
			if other.Contains(elem) {
				intersection.Add(elem)
			}
		}
	} else {
		for elem := range *o {
			if s.Contains(elem) {
				intersection.Add(elem)
			}
		}
	}
	return &intersection
}

func (s *threadUnsafeSet) Difference(other Set) Set {
	_ = other.(*threadUnsafeSet)

	difference := newThreadUnsafeSet()
	for elem := range *s {
		if !other.Contains(elem) {
			difference.Add(elem)
		}
	}
	return &difference
}

func (s *threadUnsafeSet) ToSlice() []interface{} {
	keys := make([]interface{}, 0, s.Count())
	for elem := range *s {
		keys = append(keys, elem)
	}

	return keys
}

func (s *threadUnsafeSet) Count() int {
	return len(*s)
}

func (s *threadUnsafeSet) Equal(other Set) bool {
	_ = other.(*threadUnsafeSet)

	if s.Count() != other.Count() {
		return false
	}
	for elem := range *s {
		if !other.Contains(elem) {
			return false
		}
	}
	return true
}
