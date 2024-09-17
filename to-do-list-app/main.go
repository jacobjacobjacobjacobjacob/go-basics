package main

import (
	"fmt"
)

var tasks []string // Global slice for all tasks

// Display the menu
func displayMenu() {
	fmt.Println("Hello")
}

// Add a new task
func addTask() {
	var task string
}

// List all tasks
func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("To-do list is empty.")
		return
	}
}

// Remove a task by index
func removeTask() {
	listTasks()
	var index int
}
