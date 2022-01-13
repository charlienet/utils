package set

type Set interface {
	Add(interface{}) bool
	Clear()
	Clone() Set
	Count() int
	Remove(i interface{})
	Contains(i ...interface{}) bool
	Intersect(other Set) Set
	Difference(other Set) Set
	Union(other Set) Set
	ToSlice() []interface{}
	Equal(other Set) bool
}

func NewSet(e ...interface{}) Set {
	s := newThreadSafeSet()
	for _, item := range e {
		s.Add(item)
	}

	return &s
}

func NewSetWith(es ...interface{}) Set {
	return NewSetFromSlice(es)
}

func NewSetFromSlice(s []interface{}) Set {
	a := NewSet(s...)
	return a
}

func NewThreadUnsafeSet() Set {
	set := newThreadUnsafeSet()
	return &set
}

func NewThreadUnsafeSetFromSlice(eles []interface{}) Set {
	s := NewThreadUnsafeSet()
	for _, item := range eles {
		s.Add(item)
	}
	return s
}
