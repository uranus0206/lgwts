package main

// Calculate summary of 5 numbers
func Sum(numbers []int) (sum int) {
	for _, v := range numbers {
		sum += v
	}

	return sum
}

// Calculate summarizes of any numbers array.
func SumAll(numbersToSum ...[]int) (sums []int) {
	lenOfNumbers := len(numbersToSum)
	sums = make([]int, lenOfNumbers)

	for i, v := range numbersToSum {
		sums[i] = Sum(v)
	}

	return
}
