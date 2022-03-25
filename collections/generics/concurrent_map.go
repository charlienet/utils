package generics

import (
	"runtime"
	"sync"
	"unsafe"
)

var defaultNumOfBuckets = runtime.GOMAXPROCS(runtime.NumCPU())

type concurrnetMap[K comparable, V any] struct {
	buckets      []*innerMap[K, V]
	numOfBuckets int
}

type innerMap[K comparable, V any] struct {
	lock sync.RWMutex
	m    map[K]V
}

func createInnerMap[K comparable, V any]() *innerMap[K, V] {
	return &innerMap[K, V]{
		m: make(map[K]V),
	}
}

func (im *innerMap[K, V]) get(key K) (V, bool) {
	im.lock.RLock()
	v, ok := im.m[key]
	im.lock.RUnlock()
	return v, ok
}

func (im *innerMap[K, V]) set(k K, v V) {
	im.lock.Lock()
	im.m[k] = v
	im.lock.Unlock()
}

func (im *innerMap[K, V]) del(k K) {
	im.lock.Lock()
	delete(im.m, k)
	im.lock.Unlock()
}

func (im *innerMap[K, V]) clone() *innerMap[K, V] {
	im.lock.RLock()
	new := make(map[K]V, len(im.m))
	for k, v := range im.m {
		new[k] = v
	}
	im.lock.Unlock()

	return &innerMap[K, V]{
		m: new,
	}
}

func NewConcurrnetMap[K comparable, V any]() *concurrnetMap[K, V] {
	num := defaultNumOfBuckets

	buckets := make([]*innerMap[K, V], num)
	for i := 0; i < num; i++ {
		buckets[i] = createInnerMap[K, V]()
	}

	return &concurrnetMap[K, V]{
		numOfBuckets: num,
		buckets:      buckets,
	}
}

func (m *concurrnetMap[K, V]) Set(key K, value V) {
	m.getBucket(key).set(key, value)
}

func (m *concurrnetMap[K, V]) Get(key K) (V, bool) {
	return m.getBucket(key).get(key)
}

func (m *concurrnetMap[K, V]) Delete(key K) {
	im := m.getBucket(key)
	im.del(key)
}

func (m *concurrnetMap[K, V]) Clone() *concurrnetMap[K, V] {
	buckets := make([]*innerMap[K, V], m.numOfBuckets)
	for i := 0; i < m.numOfBuckets; i++ {
		buckets[i] = m.buckets[i].clone()
	}

	return &concurrnetMap[K, V]{
		buckets:      buckets,
		numOfBuckets: m.numOfBuckets,
	}
}

func (m *concurrnetMap[K, V]) getBucket(k K) *innerMap[K, V] {
	pointer := unsafe.Pointer(&k)
	num := *(*uint)(pointer)

	id := num % uint(m.numOfBuckets)
	return m.buckets[id]
}