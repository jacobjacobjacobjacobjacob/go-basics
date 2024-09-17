// Go Course by Jon Calhoun
// Exercise #1: Quiz Game

/*
Create a program that will read in a quiz provided via a CSV file.
Quiz the user and keep track of how many questions they get correct and incorrect.
Regardless of whether the answer is correct or wrong, immediately move on to the next question.
At the end of the quiz, output total number of correct questions and how many questions there were in total.
*/

package main

import (
	"encoding/csv"
	"fmt"
	"os"
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

// Check if the answer is correct

// Main game loop
func StartGame() {
	questionData := readCsv("data/problems.csv")
	fmt.Println(questionData)
	fmt.Println(questionData[0].Question)

	// Variable to keep count of the number of questions asked
	var questionCount int
	questionCount = 1

	// Iterate over the slice and print the questions
	for _, question := range questionData {
		fmt.Printf("Q%d: %s = \n", questionCount, question.Question)
		questionCount++ // Increment count

	}

}

func main() {
	StartGame()

}
