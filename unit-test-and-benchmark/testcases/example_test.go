package unittest

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// all test cases in a unit test
func TestBiggerThenN(t *testing.T) {
	// Arrange
	tcs := []struct {
		tcName string
		num    int
		expect string
	}{
		{
			"greater then or equal to 80",
			90,
			"A",
		},
		{
			"greater then or equal to 60",
			70,
			"B",
		},
		{
			"greater then or equal to 40",
			50,
			"C",
		},
		{
			"greater then or equal to 20",
			30,
			"D",
		},
		{
			"less then 20",
			10,
			"E",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.tcName, func(t *testing.T) {
			// may need some other arranges for specific tc

			// Act
			result := showRange(tc.num)

			// Arrange
			assert.Equal(t, tc.expect, result)
		})
	}
}

// a test case in a test func
func TestBiggerThenN_ShouldReturnA_IfNumBuggerThen80(t *testing.T) {
	// Arrange
	num := 90
	expect := "A"

	// Act
	result := showRange(num)

	// Arrange
	assert.Equal(t, expect, result)
}

func TestBiggerThenN_ShouldReturnB_IfNumBuggerThen60(t *testing.T) {
	// Arrange
	num := 70
	expect := "B"

	// Act
	result := showRange(num)

	// Arrange
	assert.Equal(t, expect, result)
}

func TestBiggerThenN_ShouldReturnC_IfNumBuggerThen40(t *testing.T) {
	// Arrange
	num := 50
	expect := "C"

	// Act
	result := showRange(num)

	// Arrange
	assert.Equal(t, expect, result)
}

func TestBiggerThenN_ShouldReturnD_IfNumBuggerThen20(t *testing.T) {
	// Arrange
	num := 30
	expect := "D"

	// Act
	result := showRange(num)

	// Arrange
	assert.Equal(t, expect, result)
}

func TestBiggerThenN_ShouldReturnD_IfNumLessThen20(t *testing.T) {
	// Arrange
	num := 10
	expect := "E"

	// Act
	result := showRange(num)

	// Arrange
	assert.Equal(t, expect, result)
}
