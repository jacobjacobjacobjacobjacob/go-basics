package main

import (
	"fmt"
	"math/rand"
)

// Generate a random number between 1 and 100
func getRandomNumber() int {
	return rand.Intn(101)
}

// Take user input and check that the input is valid
func getUserGuess() int {
	var userGuess int

	for {
		fmt.Println("Input your guess (1-100):")
		fmt.Scanln(&userGuess)

		// Check if the input is within range
		if userGuess >= 1 && userGuess <= 100 {
			break
		} else {
			fmt.Println("Guess is out of range. Please enter a number between 1 and 100.")
		}
	}

	return userGuess
}

// Give hint if the number is too high or too low
func giveHint(userGuess int, randomNumber int) {
	if userGuess > randomNumber {
		fmt.Println("Your guess is too high.")
	} else {
		fmt.Println("Your guess is too low.")
	}

}

// Main game function
func StartGame() {
	randomNumber := getRandomNumber()
	fmt.Println(randomNumber) // DEBUGGING

	for {
		userGuess := getUserGuess()

		if userGuess == randomNumber {
			fmt.Println("Correct, you win!")
			break
		} else {
			fmt.Println("Wrong number.")
			giveHint(userGuess, randomNumber)
		}
	}
}

func main() {
	fmt.Println("Welcome to the guessing game!\nThe only objective is to guess a number between 1 and 100...")
	StartGame()
}
