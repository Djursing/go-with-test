package main

func Sum(nums []int) (sum int) {
	for _, n := range nums {
		sum += n
	}

	return sum
}
