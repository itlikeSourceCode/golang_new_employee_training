package main

import (
	"testing"
)

func TestUpLevel(t *testing.T) {
	UpLevel(10)

	if level != 2 || exp != 10 {
		t.Fatalf("level(%d) != 2 or exp(%d) != 10", level, exp)
	}

	UpLevel(5)

	if level != 2 || exp != 15 {
		t.Fatalf("level(%d) != 2 or exp(%d) != 15", level, exp)
	}

	UpLevel(6)

	if level != 3 || exp != 21 {
		t.Fatalf("level(%d) != 3 or exp(%d) != 21", level, exp)
	}
}
