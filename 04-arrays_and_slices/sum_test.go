package main

import (
	"reflect"
	"testing"
)

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

func TestSumAll(t *testing.T) {
	arr1 := []int{1, 2}
	arr2 := []int{0, 9}

	got := SumAll(arr1, arr2)
	exp := []int{3, 9}

	if !reflect.DeepEqual(got, exp) {
		t.Errorf("expected %d, got %d, given %v, %v", exp, got, arr1, arr2)
	}
}

func TestSumAllTails(t *testing.T) {
	arr1 := []int{1, 2}
	arr2 := []int{0, 9}

	got := SumAllTails(arr1, arr2)
	exp := []int{2, 9}

	if !reflect.DeepEqual(got, exp) {
		t.Errorf("got %q, expected %q", exp, got)
	}
}
