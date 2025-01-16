package Week1

import (
	"fmt"
)

func oddEven(n int) {
	if n%2 == 0 {
		fmt.Println("Given number is Even")
	} else {
		fmt.Println("Given number is Odd")
	}
}

func grade(score int) string {
	if score >= 90 {
		return "A"
	} else if score >= 80 {
		return "B"
	} else if score >= 70 {
		return "C"
	} else if score >= 60 {
		return "D"
	} else {
		return "F"
	}
}

func Day3() {
	var num int
	fmt.Print("Enter a number to check if it's odd or even: ")
	fmt.Scanln(&num)
	oddEven(num)

	var score int
	fmt.Print("Enter your score to get the grade:  ")
	fmt.Scanln(&score)
	yourGrade := grade(score)
	fmt.Printf("Your Grade is: %v", yourGrade)
}
