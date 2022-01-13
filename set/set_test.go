package set_test

import (
	"testing"

	"github.com/charlienet/utils/set"
)

func TestNewSet(t *testing.T) {
	ss := set.NewSet()
	ss.Add("abc")
	t.Log(ss.Contains("abc"))
	t.Log(ss.Contains("bd"))
}
