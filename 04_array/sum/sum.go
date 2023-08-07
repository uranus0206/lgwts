package main

// Calculate summary of 5 numbers
func Sum(numbers []int) (sum int) {
	for _, v := range numbers {
		sum += v
	}

	return sum
}
