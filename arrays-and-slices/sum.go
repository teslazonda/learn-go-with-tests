package main

import "fmt"

func Sum(numbers []int) int { // numbers is a array of 5 elements. Array sizes are fixed.
	sum := 0
	for _, number := range numbers { // ranges returns two values index, value. Using the blank operator to ignore indexes.
		sum += number
	}
	return sum
}

// An interesting property of arrays is that the size is encoded in its type. If you try to pass an [4]int into a function that expects [5]int, it won't compile. They are different types so it's just the same as trying to pass a string into a function that wants an int.

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)
	fmt.Println("This is the sums", sums)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}
