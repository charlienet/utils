package generics

import "sync"

type RWLockMap[K comparable, V any] struct {
	m    map[K]V
	lock sync.RWMutex
}

func NewRWLockMap[K comparable, V any]() *RWLockMap[K, V] {
	return &RWLockMap[K, V]{
		m: make(map[K]V),
	}
}

func (m *RWLockMap[K, V]) Get(key K) (V, bool) {
	m.lock.RLock()
	v, ok := m.m[key]
	m.lock.RUnlock()
	return v, ok
}

func (m *RWLockMap[K, V]) Set(key K, value V) {
	m.lock.Lock()
	m.m[key] = value
	m.lock.Unlock()
}

func (m *RWLockMap[K, V]) Delete(key K) {
	m.lock.Lock()
	delete(m.m, key)
	m.lock.Unlock()
}
