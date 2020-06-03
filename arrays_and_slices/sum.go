package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// range lets you iterate over an array
// range returns two values: index and the value
// index value is ignored above with _ (blank indentifier)
// slices are more often used in Go thank arrays as the size of collection is not encoded and can be of any size

func SumAllTails(numbersToSum ...[]int) []int {
	// lengthOfNumbers := len(numbersToSum)
	// sums := make([]int, lengthOfNumbers)
	// `make` allows slice creation with starting capacity

	// for i, numbers := range numbersToSum {
	// sums[i] = Sum(numbers)
	//}

	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}

// `append` takes a new slice and new value, and return new slice with all items in it
// variadic functions: can take variable number of args
