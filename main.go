package main

import (
	"encoding/json"
	"fer2ap/cli-go-do-it/util"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) == 1 || args[1] == "list" {
		listTasks()
	}
	if args[1] == "create" {
		createTask()
	}
	if args[1] == "deleteAll" {
		deleteAll()
	}
}

func deleteAll() {
	filePath, err := util.GetFilePath("tasks")
	if err != nil {
		panic(err)
	}
	exists, err := util.FileExists(filePath)
	if err != nil {
		panic(err)
	}
	if exists {
		err = saveTasks([]Task{})
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Tasks deleted")
}

func listTasks() {
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
	} else {
		fmt.Println("There are no tasks here. Use the create method to create your first one, i.e. gdi create")
	}
}

func createTask() {
	filePath, err := util.GetFilePath("tasks")
	if err != nil {
		panic(err)
	}
	exists, err := util.FileExists(filePath)
	if err != nil {
		panic(err)
	}
	if exists {
		fmt.Println("File recovered")
		tasks, err := recoverTasks()
		if err != nil {
			panic(err)
		}
		task := createNewTask()
		fmt.Println("Creating task: ", task)
		err = saveTasks(append(tasks, task))
		if err != nil {
			panic(err)
		}
		printMenu(append(tasks, task))
	} else {
		fmt.Println("File created")
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
		fmt.Println("Error getting file path")
		return nil, err
	}
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file")
		return nil, err
	}
	var tasks []Task
	if len(content) == 0 {
		tasks = []Task{}
	} else {
		err = json.Unmarshal(content, &tasks)
		if err != nil {
			fmt.Println("Error unmarshalling file")
			return nil, err
		}
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
	// fmt.Println(string(j))
	return nil
}

type Task struct {
	Description string
	Done        bool
}
