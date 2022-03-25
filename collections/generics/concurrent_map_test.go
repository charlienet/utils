package generics_test

import (
	"runtime"
	"testing"

	"github.com/charlienet/utils/collections/generics"
)

func TestConcurrentMap(t *testing.T) {
	t.Log(runtime.GOMAXPROCS(runtime.NumCPU()))

	key := "abc"
	value := "bcd"

	m := generics.NewConcurrnetMap[string, string]()
	m.Set(key, value)
	v, ok := m.Get(key)
	t.Log("v:", v, ok)

	m.Delete(key)
	v, ok = m.Get(key)
	t.Log("v:", v, ok)
}

func BenchmarkConcurrnetMap(b *testing.B) {
	key := "abc"
	value := "bcd"

	m := generics.NewConcurrnetMap[string, string]()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			m.Set(key, value)
			m.Get(key)
			m.Delete(key)
		}
	})
}

func BenchmarkRWLockMap(b *testing.B) {
	key := "abc"
	value := "bcd"

	m := generics.NewRWLockMap[string, string]()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			m.Set(key, value)
			m.Get(key)
			m.Delete(key)
		}
	})
}
