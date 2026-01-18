package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Priority int

const (
	uncertain Priority = iota
	low
	medium
	high
)

type Task struct {
	iscompleted bool
	desc        string
	priority    Priority
}

type List struct {
	tasks [8]Task
	count int
}

var longest_desc = 0

func (list *List) Display() {
	println()

	if list.count == 0 {
		println("To-Do list is empty")
	}

	for i := 0; i < list.count; i++ {
		fmt.Printf("%d │ %c %-*s │ %s\n",
			i+1,
			StatusToRune(list.tasks[i].iscompleted),
			longest_desc,
			list.tasks[i].desc,
			PriorityToStr(list.tasks[i].priority))
	}

	println()
}

func (list *List) Add(desc string, priority Priority) {
	list.tasks[list.count] = Task{desc: desc, priority: priority}
	list.count++
}

func (list *List) Delete(index int) {
	list.count--
	for i := 0; i < list.count; i++ {
		if i >= index {
			list.tasks[i] = list.tasks[i+1]
		}
	}
}

func (task *Task) Edit(desc string, priority Priority) {
	task.desc = desc
	task.priority = priority
}

func (task *Task) ChangeStatus() {
	task.iscompleted = !task.iscompleted
}

func (list *List) Clear() {
	var count = list.count
	for i := 1; i <= count; i++ {
		list.Delete(i)
	}
}

func (list *List) Load() {
	r, err := os.Open("todo.txt")
	if err != nil {
		return
	}
	defer r.Close()

	reader := bufio.NewReader(r)

	for {
		var task Task

		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			println("Error reading iscomplete:", err)
			break
		}
		task.iscompleted = strings.TrimSpace(line) == "true"

		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			println("Error reading desc:", err)
			break
		}
		task.desc = strings.TrimSpace(line)

		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			println("Error reading priority:", err)
			break
		}

		priorityVal, err := strconv.Atoi(strings.TrimSpace(line))
		if err != nil {
			println("Error parsing priority:", err)
			task.priority = uncertain
		} else {
			task.priority = Priority(priorityVal)
		}

		if list.count < len(list.tasks) {
			list.tasks[list.count] = task
			list.count++
		} else {
			println("Task list is full")
			break
		}

		if err == io.EOF {
			break
		}
	}
}

func (list *List) WriteToFile() {
	_ = os.Remove("todo.txt")

	var file, err = os.OpenFile("todo.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		println("Unable to create or open file:", err)
		os.Exit(1)
	}
	defer file.Close()

	for i := 0; i < list.count; i++ {
		_, err := fmt.Fprintf(file, "%t\n%s\n%d\n", list.tasks[i].iscompleted, list.tasks[i].desc, list.tasks[i].priority)
		if err != nil {
			println("Cannot to write to file:", err)
			os.Exit(1)
		}
	}
}

func StatusToRune(status bool) rune {
	if status {
		return '●'
	} else {
		return '○'
	}
}

func PriorityToStr(priority Priority) string {
	switch priority {
	case low:
		return "Low"
	case medium:
		return "Medium"
	case high:
		return "High"
	default:
		return "Uncertain"
	}
}

func FindLongestDesc(list List) {
	for i := 0; i < list.count; i++ {
		var desc_len int = len(list.tasks[i].desc)
		if desc_len > longest_desc {
			longest_desc = desc_len
		}
	}
}
