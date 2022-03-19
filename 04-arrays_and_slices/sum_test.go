package main

import "testing"

func TestSum(t *testing.T) {
	nums := [5]int{1, 2, 3, 4, 5}

	got := Sum(nums)
	exp := 15

	if got != exp {
		t.Errorf("expected %d want %d given, %v", exp, got, nums)
	}
}
