package main

import (
	"testing"
)

func increase(a, b int) int {
	return a + b
}

func TestIncrease(t *testing.T) {
	t.Log("Start testing")
	_ = increase(1, 2)
}
