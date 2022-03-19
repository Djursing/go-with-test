package main

import "testing"

func TestSum(t *testing.T) {

	t.Run("collection of any size", func(t *testing.T) {
		nums := []int{1, 2, 3}

		got := Sum(nums)
		exp := 6

		if got != exp {
			t.Errorf("expected %d, want %d, given %v", exp, got, nums)
		}
	})
}
