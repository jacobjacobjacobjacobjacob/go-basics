package main

import (
	"errors"
	"fmt"
	"math/rand"
)

// Struct for difficulty levels
type DifficultySettings struct {
	Max     int // Maximum number in the random number raneg
	Guesses int // Number of guesses allowed
}

// Generate a random number between 1 and 100
func getRandomNumber(maxNumber int) int {
	return rand.Intn(maxNumber) + 1
}

// Take user input and check that the input is valid
func getUserGuess(maxNumber int) int {
	var userGuess int

	for {
		fmt.Printf("Input your guess (1-%d):", maxNumber)
		fmt.Scanln(&userGuess)

		// Check if the input is within range
		if userGuess >= 1 && userGuess <= maxNumber {
			break
		} else {
			fmt.Printf("Guess is out of range. Please enter a number between 1 and %d.\n", maxNumber)
		}
	}

	return userGuess
}

// Give hint if the number is too high or too low
func giveHint(userGuess int, randomNumber int) {
	// Count the number of guesses
	if userGuess > randomNumber {
		fmt.Println("Your guess is too high.")
	} else {
		fmt.Println("Your guess is too low.")
	}

}

// Main game function
func StartGame(difficulty string) {
	var guesses int

	maxNumber, maxGuesses, err := getDifficultySettings(difficulty)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Start message
	fmt.Printf("Guess a number between 1 and %d.\nYou have %d available guesses.\n", maxNumber, maxGuesses)

	randomNumber := getRandomNumber(maxNumber)

	for {
		userGuess := getUserGuess(maxNumber)
		guesses++ // Increment number of guesses

		if guesses == maxGuesses {
			fmt.Println("You lost, no more guesses left.")
			fmt.Printf("The correct number was %d.\n", randomNumber)
			break
		}

		if userGuess == randomNumber {
			fmt.Printf("Correct, you win! (%d/%d guesses used).\n", guesses, maxGuesses)
			break
		} else {
			giveHint(userGuess, randomNumber)
		}
	}
}

// Check if the chosen difficulty is valid
func validateDifficulty(inputDifficulty string) (string, error) {
	switch inputDifficulty {
	case "1":
		return "easy", nil
	case "2":
		return "medium", nil
	case "3":
		return "hard", nil
	case "4":
		return "impossible", nil
	default:
		return "", errors.New("Invalid difficulty level.")
	}
}

// Fetch the parameters for the chosen difficulty settings
func getDifficultySettings(difficultyLevel string) (maxNumber int, maxGuesses int, err error) {

	// Map for difficulty settings
	difficulties := map[string]DifficultySettings{
		"easy":       {50, 10},
		"medium":     {100, 15},
		"hard":       {200, 20},
		"impossible": {300, 10},
	}

	// Fetch values from the map
	settings, exists := difficulties[difficultyLevel]
	if exists {
		maxNumber = settings.Max
		maxGuesses = settings.Guesses
	} else {
		err = errors.New("Invalid difficulty level.")
	}

	return

}

func main() {
	var inputDifficulty string
	var chosenDifficulty string
	var err error

	for {
		fmt.Println("Welcome to the guessing game!\nThe only objective is to guess the correct number.\nPlease choose a difficulty level:\n1. Easy\n2. Medium\n3. Hard\n4. Impossible")
		fmt.Scanln(&inputDifficulty)

		chosenDifficulty, err = validateDifficulty(inputDifficulty)

		if err != nil {
			fmt.Println(err)
			continue // Keep asking for input until the difficulty level is valid
		}
		break
	}

	StartGame(chosenDifficulty)
}
