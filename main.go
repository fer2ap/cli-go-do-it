package main

import (
	"encoding/json"
	"fer2ap/cli-go-do-it/util"
	"fmt"
	"os"
)

func main() {
	// task1 := Task{"Test operation 1", false}
	// task2 := Task{"Save file 2", false}
	// todolist := []Task{task1, task2}
	filePath, err := util.GetFilePath("tasks")
	if err != nil {
		panic(err)
	}
	exists, err := util.FileExists(filePath)
	if err != nil {
		panic(err)
	}
	if exists {
		tasks, err := recoverTasks()
		if err != nil {
			panic(err)
		}
		printMenu(tasks)
		task := createNewTask()
		err = saveTasks(append(tasks, task))
		if err != nil {
			panic(err)
		}
		printMenu(append(tasks, task))
	} else {
		task := createNewTask()
		tasks := []Task{task}
		err = saveTasks(tasks)
		if err != nil {
			panic(err)
		}
		printMenu(tasks)
	}
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
	filePath, err := util.GetFilePath("tasks")
	if err != nil {
		return nil, err
	}
	content, err := os.ReadFile(filePath)
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

func saveTasks(tasks []Task) error {
	j, err := json.Marshal(&tasks)
	if err != nil {
		return err
	}
	filePath, err := util.GetFilePath("tasks")
	if err != nil {
		return err
	}
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(j)
	if err != nil {
		return err
	}
	fmt.Println("Task saved")
	// fmt.Println(string(j))
	return nil
}

type Task struct {
	Description string
	Done        bool
}
