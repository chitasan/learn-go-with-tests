package iteration

func Repeat(character string, times int) string {
	var repeated string
	for i := 0; i < times; i++ {
		repeated += character
	}
	return repeated
}

// := shorthand for declaring and initializing variables; only within functions
