package task

import (
	"encoding/json"
	"fmt"
	"os"
)

func Update() {
	fmt.Println("Loading your tasks... \n\n\n")
	// Read tasks from file
	tasks, err := readTasksFromFile("tasks.json")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Tasks are shown with their IDs:\n\n")

	// Print each task
	for _, t := range tasks {
		fmt.Printf("%d  >>> %s\n\n", t.ID, t.Sentence)

	}

	var task_inp int
	fmt.Println("Please Enter ID of task you wish to edit or delete")
C:

	_, ok := fmt.Scanln(&task_inp)
	if ok != nil {
		fmt.Println("Please enter correct input")
		goto C
	}
	exist := false
	var index int
	for i, t := range tasks {
		if t.ID == task_inp {
			exist = true
			index = i
			break
		}
	}
	if !exist {
		fmt.Println("Please Specify correct ID")
		goto C
	}
	fmt.Println("Task found with ID", task_inp)

	TaskAction(index, tasks)
}

func TaskAction(i int, t []task) {
	fmt.Println("What action you want to perform?")
	fmt.Println("1: Edit")
	fmt.Println("2: Delete")

	var task_act int
D:
	_, ok := fmt.Scanln(&task_act)
	if ok != nil || task_act > 2 || task_act < 1 {
		fmt.Println("Please enter correct input")
		goto D
	}

	switch task_act {
	case 1:
		EditData(i, t)
	case 2:
		DeleteData(i, t)
	}

}
func EditData(index int, tasks []task) {
	fmt.Println("Editing your task, press Enter to save the editted version")
	user_edit := getUserInput("Enter your todo task: ")
	tasks[index].Sentence = user_edit
	err := writeTasksToFile("tasks.json", tasks)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Task Edited successfully.")
}

func DeleteData(index int, tasks []task) {
	fmt.Println("Are you sure want to delete?")
	fmt.Printf("%s \n\n", tasks[index].Sentence)
	fmt.Println("Type `yes` to confirm, anything else to cancel")
	var user_delete string
	_, ok := fmt.Scanln(&user_delete)
	if ok != nil || user_delete != "yes" {
		fmt.Println("Delete Operation Cancelled.")
		return
	}
	if user_delete == "yes" {
		tasks = append(tasks[:index], tasks[index+1:]...)
		err := writeTasksToFile("tasks.json", tasks)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Task deleted successfully.")
	}

}

// Function to write tasks to file
func writeTasksToFile(filename string, tasks []task) error {
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
