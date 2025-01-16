package Week1

import (
	"fmt"
)

func swap(x *int, y *int) {
	temp := *x
	*x = *y
	*y = temp
}

func Day2() {
	var (
		name string = "Basit"
		age  int    = 21
		city string = "Gandhinagar"
	)
	fmt.Printf("I'm %v, I'm %v years old and I live in %v \n", name, age, city)

	a := 10
	b := 20
	swap(&a, &b)
	fmt.Println(a, b)
}
