package iteration

// Return repeated string with desired times.
func Repeat(element string, times int) string {
	var repeated string

	for i := 0; i < times; i++ {
		repeated += element
	}

	return repeated
}
