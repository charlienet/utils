package rand_test

import (
	"encoding/hex"
	"testing"

	"github.com/charlienet/utils/rand"
)

func TestRandString(t *testing.T) {
	t.Log(rand.AllChars.RandString(10))

	b, err := rand.RandByts(32)
	t.Log(err)
	t.Log(hex.EncodeToString(b))
}
