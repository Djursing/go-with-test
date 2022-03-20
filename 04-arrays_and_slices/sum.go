package main

func Sum(nums []int) (sum int) {
	for _, n := range nums {
		sum += n
	}

	return
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, n := range numbersToSum {
		sums = append(sums, Sum(n))
	}

	return
}

func SumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, n := range numbersToSum {
		if len(n) == 0 {
			sums = append(sums, 0)
			continue
		}

		tail := n[1:]
		sums = append(sums, Sum(tail))
	}

	return
}
