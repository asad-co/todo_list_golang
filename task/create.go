package task

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func Create() {

	sentence := getUserInput("Enter your todo task: ")

	t := task{sentence, getTime(), generateRandomID()}
	err := appendTaskToFile("tasks.json", t)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Success!")
}
func getTime() string {
	// Get current time
	currentTime := time.Now()

	// Define the layout for formatting the time
	layout := "2nd January 2006 15:04:05"

	// Format the time according to the layout
	formattedTime := currentTime.Format(layout)

	return formattedTime
}

type task struct {
	Sentence string `json:"sentence"`
	Time     string `json:"time"`
	ID       int    `json:"id"`
}

// Function to append a new task to the existing file or create a new file if it doesn't exist
func appendTaskToFile(filename string, t task) error {
	// Read existing tasks from file
	tasks, err := readTasksFromFile(filename)
	if err != nil {
		return err
	}

	// Append the new task to the existing tasks
	tasks = append(tasks, t)

	// Convert tasks slice to JSON
	data, err := json.MarshalIndent(tasks, "", "    ")
	if err != nil {
		return err
	}

	// Write JSON data to file
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

// Function to read tasks from file
func readTasksFromFile(filename string) ([]task, error) {
	var tasks []task

	// Check if file exists
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		// If file doesn't exist, return an empty slice
		return tasks, nil
	}

	// Read file contents
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON data into tasks slice
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		return scanner.Text()
	}
	return ""
}

// Function to generate a random ID
func generateRandomID() int {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate a random number between 1000 and 9999
	return rand.Intn(900000) + 100000
}
