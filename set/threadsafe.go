package set

import "sync"

type threadSafeSet struct {
	s    threadUnsafeSet
	lock sync.RWMutex
}

func newThreadSafeSet() threadSafeSet {
	return threadSafeSet{s: newThreadUnsafeSet()}
}

func (s *threadSafeSet) Add(i interface{}) bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	ret := s.s.Add(i)

	return ret
}

func (set *threadSafeSet) Contains(i ...interface{}) bool {
	set.lock.RLock()
	defer set.lock.RUnlock()

	ret := set.s.Contains(i...)
	return ret
}

func (set *threadSafeSet) Remove(i interface{}) {
	set.lock.RLock()
	defer set.lock.RUnlock()

	delete(set.s, i)
}

func (s *threadSafeSet) Clear() {
	s.s = newThreadUnsafeSet()
}

func (s *threadSafeSet) Clone() Set {
	s.lock.RLock()
	defer s.lock.RUnlock()

	unsafeClone := s.s.Clone().(*threadUnsafeSet)
	ret := &threadSafeSet{s: *unsafeClone}

	return ret
}

func (s *threadSafeSet) Intersect(other Set) Set {
	o := other.(*threadSafeSet)

	s.lock.RLock()
	o.lock.RLock()

	unsafeIntersection := s.s.Intersect(&o.s).(*threadUnsafeSet)
	ret := &threadSafeSet{s: *unsafeIntersection}
	s.lock.RUnlock()
	o.lock.RUnlock()
	return ret
}

func (s *threadSafeSet) Union(other Set) Set {
	o := other.(*threadSafeSet)

	s.lock.RLock()
	o.lock.RLock()

	unsafeUnion := s.s.Union(&o.s).(*threadUnsafeSet)
	ret := &threadSafeSet{s: *unsafeUnion}
	s.lock.RUnlock()
	o.lock.RUnlock()
	return ret
}

func (s *threadSafeSet) Difference(other Set) Set {
	o := other.(*threadSafeSet)

	s.lock.RLock()
	o.lock.RLock()

	unsafeDifference := s.s.Difference(&o.s).(*threadUnsafeSet)
	ret := &threadSafeSet{s: *unsafeDifference}
	s.lock.RUnlock()
	o.lock.RUnlock()
	return ret
}

func (s *threadSafeSet) ToSlice() []interface{} {
	keys := make([]interface{}, 0, s.Count())
	s.lock.RLock()
	for elem := range s.s {
		keys = append(keys, elem)
	}
	s.lock.RUnlock()
	return keys
}

func (s *threadSafeSet) Count() int {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return len(s.s)
}

func (s *threadSafeSet) Equal(other Set) bool {
	o := other.(*threadSafeSet)

	s.lock.RLock()
	o.lock.RLock()

	ret := s.s.Equal(&o.s)
	s.lock.RUnlock()
	o.lock.RUnlock()
	return ret
}
