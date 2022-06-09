package unittest

import (
	"testing"
)

func TestLongtimeWithParallel(t *testing.T) {
	tcs := []struct {
		tcName string
	}{
		{
			"t1",
		},
		{
			"t2",
		},
		{
			"t3",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.tcName, func(t *testing.T) {
			t.Parallel()
			longtime()
		})
	}
}
