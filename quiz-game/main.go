package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

// Struct to store the CSV data
type Record struct {
	Question string
	Answer   string
}

// Read the CSV file
func readCsv(csvFilePath string) []Record {
	file, err := os.Open(csvFilePath) // Open file

	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	defer file.Close() // Close

	reader := csv.NewReader(file)    // Create CSV reader
	records, err := reader.ReadAll() // Read all records

	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	// Slice to hold the parsed records
	var data []Record

	// Iterate over the records and store them in the slice
	for _, record := range records {
		r := Record{
			Question: record[0],
			Answer:   record[1],
		}

		data = append(data, r)
	}

	return data

}

// Main game loop
func StartGame() {
	questionData := readCsv("data/problems.csv")

	// Variable to keep count of the number of questions and score.
	questionCount := 1
	scoreCount := 0

	// Set the timer duration (e.g., 30 seconds)
	timerDuration := 30 * time.Second
	done := make(chan bool)

	// Start the timer in a goroutine
	go func() {
		time.Sleep(timerDuration)
		done <- true // Signal that the time is up
	}()

	// Iterate over the slice and print the questions
	for _, question := range questionData {
		select {
		case <-done:
			fmt.Println("\nTime's up! The quiz has ended.")
			fmt.Printf("You scored %d out of %d.\n", scoreCount, len(questionData))
			return
		default:
			fmt.Printf("Question %d:\n%s=", questionCount, question.Question)
			questionCount++ // Increment count

			// Ask for user input
			var userAnswer string
			fmt.Scanln(&userAnswer)

			// Validate user input against the correct answer
			if userAnswer == question.Answer {
				scoreCount++
			}
		}
	}

	fmt.Printf("You scored %d out of %d.\n", scoreCount, len(questionData))
}

func main() {
	fmt.Println("Welcome to the quiz game.\nPress Enter to start...")

	// Wait for user to press Enter
	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadString('\n')
	StartGame()

}
