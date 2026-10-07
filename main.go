package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

func loadTasks() ([]Task, error) {
	data, err := os.ReadFile("tasks.json")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Файла не существует. Создание файла")
			err = os.WriteFile("tasks.json", []byte("[]"), 0644)
			if err != nil {
				return nil, err
			}
			return []Task{}, nil
		} else {
			return nil, fmt.Errorf("не удалось прочитать tasks.json: %w", err)
		}

	}
	var tasks []Task
	if len(data) == 0 {
		return []Task{}, nil
	}
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, fmt.Errorf("не удалось разобрать tasks.json: %w", err)
	}
	return tasks, nil

}
func main() {

	tasks, err := loadTasks()
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(os.Args) < 2 {
		fmt.Println("Введено неверное количество аргументов")
		return
	}
	command := os.Args[1]
	switch command {
	case "add":
		if len(os.Args) != 3 {
			fmt.Println("Введено неверное количество элементов")
			return
		}
		AddTask(&tasks, os.Args[2])

	case "update":
		if len(os.Args) != 4 {
			fmt.Println("Введено неверное количество элементов")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Ошибка перевода числа")
			return
		}
		UpdateTask(&tasks, id, os.Args[3])
	case "list":
		var status string
		if len(os.Args) > 3 {
			fmt.Println("Введено неверное количество элементов")
			return
		}
		if len(os.Args) == 2 {
			status = ""
		} else {
			status = os.Args[2]
		}
		List(tasks, status)
	case "delete":
		if len(os.Args) != 3 {
			fmt.Println("Введено неверное количество элементов")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Ошибка перевода числа")
			return
		}
		DeleteTask(&tasks, id)
	case "mark-in-progress":
		if len(os.Args) != 3 {
			fmt.Println("Введено неверное количество элементов")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Ошибка перевода числа")
			return
		}
		MarkTask(&tasks, id, "in-progress")
	case "mark-done":
		if len(os.Args) != 3 {
			fmt.Println("Введено неверное количество элементов")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Ошибка перевода числа")
			return
		}
		MarkTask(&tasks, id, "done")
	case "mark-todo":
		if len(os.Args) != 3 {
			fmt.Println("Введено неверное количество элементов")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Ошибка перевода числа")
			return
		}
		MarkTask(&tasks, id, "todo")
	default:
		fmt.Println("Неизвестная команда")
	}

}
