package main

import (
	"fmt"
	"todo/task"
)

func main() {

	var user_in int
	fmt.Println("Welcome to ToDo Application")
	fmt.Println("You can do interact here")
	fmt.Println("1: Create a todo task")
	fmt.Println("2: View your task list")
	fmt.Println("3: Edit/Delete the task")
A:
	_, ok := fmt.Scanln(&user_in)
	if ok != nil || user_in > 3 || user_in < 1 {
		fmt.Println("Please enter correct input")
		goto A
	}
	runEvent(user_in)
}

func runEvent(input_user int) {
	switch input_user {
	case 1:
		task.Create()
	case 2:
		task.View()
	case 3:
		task.Update()
	}
}
