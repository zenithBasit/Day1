package main

import (
	"fmt"
	"math"
)

func avgSum(arr []int) {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	avg := sum / len(arr)
	fmt.Printf("Sum: %v and Average: %v \n", sum, avg)
}

func minMax(a []int) {
	min := math.MaxInt
	max := math.MinInt
	for i := 0; i < len(a); i++ {
		if max < a[i] {
			max = a[i]
		}
		if min > a[i] {
			min = a[i]
		}
	}
	fmt.Printf("Maximum in slice is: %v and Minimum in slice is: %v \n", max, min)
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	avgSum(arr)

	s1 := []int{3, 4, 6, 7, 19}
	minMax(s1)

	students := make(map[string]string)

	students["Alice"] = "A"
	students["Bob"] = "B"
	students["Charlie"] = "A"
	students["Diana"] = "C"

	stu := "Bob"
	grade, exists := students[stu]
	if exists {
		fmt.Printf("\n%s's grade: %s\n", stu, grade)
	} else {
		fmt.Printf("\nNo grade found for %s.\n", stu)
	}
}
