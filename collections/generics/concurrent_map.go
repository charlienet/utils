package generics

import (
	"fmt"
	"runtime"
	"sync"

	"github.com/charlienet/utils/hash"
)

var defaultNumOfBuckets = runtime.GOMAXPROCS(runtime.NumCPU())

type ConcurrnetMap[K comparable, V any] struct {
	buckets      []*innerMap[K, V]
	numOfBuckets uint64
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

func (im *innerMap[K, V]) foreach(f func(K, V)) {
	for k, v := range im.m {
		f(k, v)
	}
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

func NewConcurrnetMap[K comparable, V any]() *ConcurrnetMap[K, V] {
	num := defaultNumOfBuckets

	buckets := make([]*innerMap[K, V], num)
	for i := 0; i < num; i++ {
		buckets[i] = createInnerMap[K, V]()
	}

	return &ConcurrnetMap[K, V]{
		numOfBuckets: uint64(num),
		buckets:      buckets,
	}
}

func (m *ConcurrnetMap[K, V]) Set(key K, value V) {
	m.getBucket(key).set(key, value)
}

func (m *ConcurrnetMap[K, V]) Get(key K) (V, bool) {
	return m.getBucket(key).get(key)
}

func (m *ConcurrnetMap[K, V]) Delete(key K) {
	im := m.getBucket(key)
	im.del(key)
}

func (m *ConcurrnetMap[K, V]) ForEach(f func(K, V)) {
	var wg sync.WaitGroup

	num := int(m.numOfBuckets)

	wg.Add(int(m.numOfBuckets))
	for i := 0; i < num; i++ {
		go func(i int) {
			m.buckets[i].foreach(f)
			wg.Done()
		}(i)
	}

	wg.Wait()
}

func (m *ConcurrnetMap[K, V]) Clone() *ConcurrnetMap[K, V] {
	num := int(m.numOfBuckets)

	buckets := make([]*innerMap[K, V], m.numOfBuckets)
	for i := 0; i < num; i++ {
		buckets[i] = m.buckets[i].clone()
	}

	return &ConcurrnetMap[K, V]{
		buckets:      buckets,
		numOfBuckets: m.numOfBuckets,
	}
}

func (m *ConcurrnetMap[K, V]) getBucket(k K) *innerMap[K, V] {
	bytes := getBytes(k)

	id := hash.XXHashUint64(bytes) % m.numOfBuckets
	return m.buckets[id]
}

func getBytes(k any) []byte {
	return []byte(fmt.Sprintf("%v", k))
}
