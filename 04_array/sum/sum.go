package main

// Calculate summary of 5 numbers
func Sum(numbers [5]int) (sum int) {
	for i := 0; i < 5; i++ {
		sum += numbers[i]
	}
	return sum
}