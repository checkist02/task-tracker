package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

var Id = 0

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
func saveToFile(tasks []Task) error{
	data, err := json.Marshal(tasks)
	if err != nil {
		return errors.New("Ошибка перевода в json")
	} else {
		err = os.WriteFile("tasks.json", data, 0644)
		if err != nil {
			return errors.New("Ошибка записи в файл")
		}
	}
	return nil
}

func printTask(task Task) {
	fmt.Printf(`{
	ID: %d
	Description: %s
	Status: %s
	CreatedAt: %s
	UpdatedAt: %s
}`, task.ID, task.Description, task.Status, task.CreatedAt.Format("02.01.2006 15:04:05"), task.UpdatedAt.Format("02.01.2006 15:04:05"))
}
func AddTask(tasks *[]Task, description string) () {
	newTask := Task{
		ID:          (getId(*tasks)),
		Description: description,
		Status:      "in-progress",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	*tasks = append(*tasks, newTask)
	
	if err := saveToFile(*tasks); err!=nil{
		fmt.Println("Ошибка добавления задачи")
		return
	}
	fmt.Println("Добавлена новая запись с id", newTask.ID)
}
func DeleteTask(tasks *[]Task, id int) {
	_, idFound, err := foundById(tasks, id)
	if err != nil {
		fmt.Println(err)
	} else {
		*tasks = append((*tasks)[:idFound], (*tasks)[idFound+1:]...)
		if err := saveToFile(*tasks); err!=nil{
			fmt.Println("Ошибка удаления задачи")
		return
		}
	}
}
func foundById(tasks *[]Task, id int) (*Task, int, error) {
	for i := range *tasks {
		if (*tasks)[i].ID == id {
			return &(*tasks)[i], i, nil
		}
	}
	return nil, 0, errors.New("Данный id не найден")
}
func UpdateTask(tasks *[]Task, id int, description string) {
	_, idFound, err := foundById(tasks, id)
	if err != nil {
		fmt.Println(err)
		return
	}

	(*tasks)[idFound].Description = description
	(*tasks)[idFound].UpdatedAt = time.Now()

	if err := saveToFile(*tasks); err != nil {
		fmt.Println("Ошибка обновления задачи")
		return
	}

	fmt.Println("Задача обновлена")
}
func List(tasks []Task, status string) {
	switch status {
	case "done":
		for i := range tasks {
			if tasks[i].Status == "done" {
				printTask(tasks[i])
			}
		}
	case "todo":
		for i := range tasks {
			if tasks[i].Status == "todo" {
				printTask(tasks[i])
			}
		}
	case "in-progress":
		for i := range tasks {
			if tasks[i].Status == "in-progress" {
				printTask(tasks[i])
			}
		}
	case "":
		for i := range tasks {
			printTask(tasks[i])
		}
	default:
		fmt.Println("Введен неверный фильтр")
	}
}
func MarkTask(tasks *[]Task, id int, status string) {
	_, i, err := foundById(tasks, id)
	if err != nil {
		fmt.Println(err)
		return
	}

	(*tasks)[i].Status = status
	(*tasks)[i].UpdatedAt = time.Now()

	if err := saveToFile(*tasks); err != nil {
		fmt.Println("Ошибка сохранения задачи")
		return
	}

	fmt.Println("Статус задачи обновлён")
}
func getId(tasks []Task) int {
	maxID := 0

	for _, task := range tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}

	return maxID + 1
}
