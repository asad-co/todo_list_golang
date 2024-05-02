package task

import (
	"encoding/json"
	"fmt"
	"os"
)

func View() {
	fmt.Println("Loading your tasks... \n\n\n")
	// Read tasks from file
	tasks, err := readTasksFromFile("tasks.json")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print each task
	for _, t := range tasks {
		fmt.Printf("%s  >>> saved at %s\n\n", t.Sentence, t.Time)

	}
}

// Function to read tasks from file
func viewTasksFromFile(filename string) ([]task, error) {
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
