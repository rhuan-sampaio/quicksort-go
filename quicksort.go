package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := os.Args[1:]
	numbers := make([]int, len(input))
	for i, v := range input {
		number, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("Invalid number")
			return
		}
		numbers[i] = number
	}
	fmt.Println(quicksort(numbers))

}

func quicksort(numbers []int) []int {
	if len(numbers) <= 1 {
		return numbers
	}
	n := make([]int, len(numbers))
	copy(n, numbers)

	pivotIndex := len(numbers) / 2
	pivot := n[pivotIndex]
	n = append(n[:pivotIndex], n[pivotIndex+1:]...)

	minors, biggers := partitionate(n, pivot)
	return append(append(quicksort(minors), pivot), quicksort(biggers)...)
}

func partitionate(numbers []int, pivot int) (minors []int, biggers []int) {
	for _, v := range numbers {
		if pivot >= v {
			minors = append(minors, v)
		} else {
			biggers = append(biggers, v)
		}
	}

	return minors, biggers
}
