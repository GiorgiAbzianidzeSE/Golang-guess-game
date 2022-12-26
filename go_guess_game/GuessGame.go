package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	var name string
	var randomeNumber = rand.Intn(100)
	chanceCounter := 1
	var inputedNumber int
	fmt.Print("Please enter your name: ")
	fmt.Scan(&name)
	fmt.Printf("Hello %s ,You have to guess the number between 0 - 100. You only have 3 chances \n", name)
	for {
		fmt.Print("input number: ")
		fmt.Scan(&inputedNumber)
		if inputedNumber == randomeNumber {
			fmt.Printf("You won in %d chance %d", chanceCounter, chanceCounter)
			return
		} else if inputedNumber > randomeNumber {
			fmt.Println("Less")
		} else {
			fmt.Println("More")
		}

		if chanceCounter == 3 {
			fmt.Println("You lost number was: ", randomeNumber)
			return
		}
		chanceCounter++
	}
}
