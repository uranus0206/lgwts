package iteration

const repeatCount = 5

func Repeat(element string) string {
	var repeated string

	for i := 0; i < repeatCount; i++ {
		repeated += element
	}

	return repeated
}
