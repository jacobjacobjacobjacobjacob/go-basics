package main

import (
	"fmt"
	"math/rand"
)

// Generate random number
func getRandomNumber() int {
	return rand.Intn(100)
}

// Main game function
func StartGame() {
	random_number := getRandomNumber()
	fmt.Println(random_number)

}

func main() {
	fmt.Println("Welcome to the guessing game!\nThe only objective is to guess a number between 1 and 100...")
	StartGame()
}
