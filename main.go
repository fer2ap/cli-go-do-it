package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	// task1 := Task{"Test operation 1", false}
	// task2 := Task{"Save file 2", false}
	// todolist := []Task{task1, task2}
	tasks, err := recoverTasks()
	if err != nil {
		panic(err)
	}
	printMenu(tasks)
	task := createNewTask()
	saveTasks(append(tasks, task))
}

func createNewTask() Task {
	fmt.Println("Write new task description (Must use quotes around input, i.e. \"New task desc.\"): ")
	var taskDescription string
	fmt.Scanf("%q", &taskDescription)
	task := Task{taskDescription, false}
	return task
}

func printMenu(tasks []Task) {
	for i, v := range tasks {
		var checkerIcon string
		if v.Done {
			checkerIcon = " OK"
		} else {
			checkerIcon = " --"
		}
		fmt.Println(" ", i, " ", "> ", v.Description, checkerIcon)
	}
}

func recoverTasks() ([]Task, error) {
	content, err := os.ReadFile("/tmp/tasks")
	if err != nil {
		return nil, err
	}
	var tasks []Task
	err = json.Unmarshal(content, &tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func saveTasks(tasks []Task) {
	j, err := json.Marshal(&tasks)
	if err != nil {
		panic(err)
	}
	f, err := os.Create("/tmp/tasks")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = f.Write(j)
	if err != nil {
		panic(err)
	}
	fmt.Println("Tasks saved:")
	fmt.Println(string(j))
}

type Task struct {
	Description string
	Done        bool
}
