package Week1

import (
	"fmt"
	"math/rand"
	// "time"
)

func guess() string {
	var num int
	// rand.Seed(time.Now().UnixNano())
	random := rand.Intn(101)
	fmt.Printf("Guess a Number: ")
	for {
		fmt.Scanln(&num)
		if num == random {
			return "Congrats, you guessed the right number"
		} else if num < random {
			fmt.Println("Guess Higher")
		} else {
			fmt.Println("Guess Lower")
		}
	}
}

func Day5() {
	fmt.Print("Welcome to Guessing Game!!!\nGuess Number between 0-100 \n")
	res := guess()
	fmt.Println(res)
}
