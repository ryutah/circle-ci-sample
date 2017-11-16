package sample

import (
	"testing"
)

func TestHello(t *testing.T) {
	if got := hello(); got != "Hello" {
		t.Errorf("want : %v, got : %v", "Hello", got)
	}
}
