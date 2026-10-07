package main

import (
	
	"fmt"
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
func printTask(task Task){
	fmt.Printf(`{
	ID: %d
	Description: %s
	Status: %s
	CreatedAt: %s
	UpdatedAt: %s
}`, task.ID, task.Description, task.Status, task.CreatedAt.Format("02.01.2006 15:04:05"), task.UpdatedAt.Format("02.01.2006 15:04:05"))
}
func AddTask(tasks *[]Task, description string) {
	newTask := Task{
		ID:          getId(*tasks),
		Description: description,
		Status:      "in-progress",
		CreatedAt:   time.Now(),
		UpdatedAt: time.Now(),
	}
	*tasks = append(*tasks, newTask)

}
func DeleteTask(tasks *[]Task, id int) {
	_, idFound := foundById(tasks, id)
	*tasks = append((*tasks)[:idFound], (*tasks)[idFound+1:]...)
}
func foundById(tasks *[]Task, id int) (*Task, int) {
	for i := range *tasks {
		if (*tasks)[i].ID == id {
			return &(*tasks)[i], i
		}
	}
	return nil, 0
}
func UpdateTask(tasks *[]Task, id int, description string) {
	_, i := foundById(tasks, id)
	(*tasks)[i].Description = description
}
func List(tasks []Task, status string){
	switch status{
	case "done":
		for i:= range tasks{
			if tasks[i].Status == "done"{
				printTask(tasks[i])
			}
		}
	case "todo":
		for i:= range tasks{
			if tasks[i].Status == "todo"{
				printTask(tasks[i])
			}
		}
	case "in-progress":
		for i:= range tasks{
			if tasks[i].Status == "in-progress"{
				printTask(tasks[i])
			}
		}
	case "":
		for i:= range tasks{
			printTask(tasks[i])
		}
	default:
		fmt.Println("Введен неверный фильтр")
	}
}
func MarkTask(tasks *[]Task, id int, status string){
	_, i := foundById(tasks, id)
	(*tasks)[i].Status = status
}
func getId(tasks []Task)(int){
	return len(tasks)+1
}
