package main



func main() {
	var tasks []Task
	AddTask(&tasks, "abd")
	UpdateTask(&tasks, 0, "ASD")
	MarkTask(&tasks, 1, "todo")
	List(tasks, "todo")
}
