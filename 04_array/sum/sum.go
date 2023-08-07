package main

// Calculate summary of 5 numbers
func Sum(numbers []int) (sum int) {
	for _, v := range numbers {
		sum += v
	}

	return sum
}

// Calculate summarizes of any numbers array.
func SumAllTails(numbersToSum ...[]int) []int {
	sums := make([]int, 0)

	for _, v := range numbersToSum {
		if len(v) == 0 {
			sums = append(sums, 0)
		} else {
			tail := v[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
