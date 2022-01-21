package rand_test

import (
	"bytes"
	"encoding/hex"
	"testing"
	"time"

	mrnd "math/rand"

	"github.com/charlienet/utils/rand"
)

func TestRandString(t *testing.T) {
	t.Log(rand.Hex.RandString(2000))

	b, err := rand.RandBytes(32)
	t.Log(err)
	t.Log(hex.EncodeToString(b))
}

func TestRandHex(t *testing.T) {
	h := rand.Hex.RandString(3)
	t.Log(h)
}

func TestRandMax(t *testing.T) {
	mrnd.Seed(time.Now().UnixNano())
}

func BenchmarkNoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
	}
}

func BenchmarkRandString(b *testing.B) {

	for i := 0; i < 10; i++ {
		rand.Hex.RandString(10)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rand.Hex.RandString(20)
	}

	// b.Run("randString", func(b *testing.B) {

	// 	for i := 0; i < b.N; i++ {
	// 		rand.Hex.RandString(256)
	// 	}
	// })

	// for i := 0; i < b.N; i++ {
	// 	rand.RandBytes(16)
	// }
}

func BenchmarkString(b *testing.B) {
	elems := []byte("abcdefghijk")

	b.Run("1", func(b *testing.B) {
		a := []byte{}
		for i := 0; i < b.N; i++ {
			for _, elem := range elems {
				a = append(a, elem)
			}
		}
	})

	b.Run("2", func(b *testing.B) {
		a := make([]byte, len(elems))
		for i := 0; i < b.N; i++ {
			for _, elem := range elems {
				a = append(a, elem)
			}
		}
	})

	b.Run("3", func(b *testing.B) {
		a := make([]byte, len(elems))
		for i := 0; i < b.N; i++ {
			a = append(a, elems...)
		}
	})
}

func BenchmarkConcatString(b *testing.B) {
	elems := []string{"1", "2", "3", "4", "5"}

	b.Run("add", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ret := ""
			for _, elem := range elems {
				ret += elem
			}
		}
	})

	b.Run("buffer", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var buf bytes.Buffer
			for _, elem := range elems {
				buf.WriteString(elem)
			}
		}
	})
}
