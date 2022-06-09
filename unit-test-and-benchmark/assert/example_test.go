package unittest

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBiggerThenN(t *testing.T) {
	// Arrange or Given
	target, n := 10, 0

	// Act or When
	result := biggerThenN(target, n)

	// Assert or Then
	assert.Equal(t, true, result)
}
