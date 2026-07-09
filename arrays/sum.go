package arrays

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numsToSum ...[]int) []int {
	var sums []int
	for _, nums := range numsToSum {
		sums = append(sums, Sum(nums))
	}

	return sums
}

func SumAllTails(numsToSum ...[]int) []int {
	var sums []int
	for _, nums := range numsToSum {
		if len(nums) < 1 {
			sums = append(sums, 0)
			continue
		}

		tail := nums[1:]
		sums = append(sums, Sum(tail))
	}
	return sums
}
