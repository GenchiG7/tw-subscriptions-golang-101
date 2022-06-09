package unittest

import (
	"testing"
)

func TestGreeting_Success(t *testing.T) {
	result := greeting("genchi")

	if result != "hello, genchi!" {
		t.Errorf("invalid freeting")
	}
}

func TestGreeting_Failed(t *testing.T) {
	result := greeting("genchi")

	if result != "not greeting string" {
		t.Errorf("invalid freeting")
	}
}
