package dateconv_test

import (
	"testing"

	"github.com/charlienet/utils/dateconv"
)

func TestToday(t *testing.T) {
	today := dateconv.Today()
	t.Log(dateconv.TimeToString(&today))
}
