package Week1

import (
	"fmt"
)

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func reverse(s string) string {
	var res string = ""
	for i := len(s) - 1; i >= 0; i-- {
		res += string(s[i])
	}
	return res
}

func Day4() {
	num := fact(7)
	fmt.Println("Factorial is: ", num)

	str := "zenithive"
	revStr := reverse(str)
	fmt.Println("Reversed string is: ", revStr)
}
