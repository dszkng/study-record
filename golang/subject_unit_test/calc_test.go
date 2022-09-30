package main

import "testing"

func TestAdd(t *testing.T) {
	if ans := add(1, 2); ans != 3 {
		t.Error("add(1, 2) should be equal to 3")
	}

	if ans := add(-2, 3); ans != 1 {
		t.Error("add(-2, 3) should be equal to 1")
	}
}

func TestMul(t *testing.T) {
	if ans := mul(1, 1); ans != 1 {
		t.Error("mul(1, 1) should be equal to 1")
	}
}
