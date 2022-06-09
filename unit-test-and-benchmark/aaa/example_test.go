package unittest

import (
	"testing"
)

func TestBiggerThenN(t *testing.T) {
	// Arrange or Given
	target, n := 10, 0

	// Act or When
	result := biggerThenN(target, n)

	// Assert or Then
	if !result {
		t.Errorf("target should bigger then n")
	}
}
