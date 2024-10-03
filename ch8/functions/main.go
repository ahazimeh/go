package main

import "fmt"

func main() {
	numbers := []int{1, 15, 10}
	sum := sumup(1, numbers...)

	fmt.Println(sum)
}

func sumup(startingValue int, numbers ...int) int {
	println(startingValue)
	sum := 0

	for _, val := range numbers {
		sum += val
	}
	return sum
}
