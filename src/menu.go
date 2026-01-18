package main

import (
	"bufio"
	"os"
	"strconv"
)

type Menu struct {
	input   string
	scanner *bufio.Scanner
}

func (menu *Menu) Draw(list List) {
	println()
	println("(a)dd new task")
	println("(d)elete task")
	println("(e)dit task")
	println("(c)hange status")
	println("(C)lear list")
	println("(q)uit")

	FindLongestDesc(list)
	list.Display()
}

func (menu *Menu) HandleInput(list *List) {
	print(">> ")
	scanner := bufio.NewScanner(os.Stdin)
	if ok := scanner.Scan(); !ok {
		println("\nIncorrect input")
		return
	}

	menu.input = scanner.Text()

	switch menu.input {
	case "a":
		menu.AddTask(list)
	case "d":
		menu.DeleteTask(list)
	case "e":
		menu.EditTask(list)
	case "c":
		menu.ChangeTaskStatus(list)
	case "C":
		list.Clear()
	}

	list.WriteToFile()
}

func (menu *Menu) AddTask(list *List) {
	print("\nDescription (64 max): ")

	menu.scanner = bufio.NewScanner(os.Stdin)
	if ok := menu.scanner.Scan(); !ok {
		println("\nIncorrect input")
		return
	}
	var desc = menu.scanner.Text()

	println("\nPriorities:")
	println("(l)ow")
	println("(m)edium")
	println("(h)igh")

	print(">> ")
	menu.scanner = bufio.NewScanner(os.Stdin)
	if ok := menu.scanner.Scan(); !ok {
		println("\nIncorrect input")
		return
	}
	menu.input = menu.scanner.Text()

	var priority Priority
	switch menu.input {
	case "l":
		priority = low
	case "m":
		priority = medium
	case "h":
		priority = high
	}

	if list.count == len(list.tasks) {
		println("\nYour list is full")
		return
	}
	if len(desc) > 64 {
		println("\nDescription cannot be longer than 64 characters")
		return
	}

	list.Add(desc, priority)
	list.WriteToFile()
}

func (menu *Menu) DeleteTask(list *List) {
	print("\nIndex: ")
	menu.scanner = bufio.NewScanner(os.Stdin)
	if ok := menu.scanner.Scan(); !ok {
		println("\nIncorrect input")
		return
	}

	var index, err = strconv.Atoi(menu.scanner.Text())
	if err != nil {
		println("\nIncorrect index", err)
		return
	}

	if index <= 0 || index > list.count {
		println("\nIndex cannot be lower than 0 or higher than task count")
		return
	}

	list.Delete(index)
	list.WriteToFile()
}

func (menu *Menu) EditTask(list *List) {

	print("\nIndex: ")
	menu.scanner = bufio.NewScanner(os.Stdin)
	if ok := menu.scanner.Scan(); !ok {
		println("\nIncorrect input")
		return
	}
	var index, err = strconv.Atoi(menu.scanner.Text())
	if err != nil {
		println("\nIncorrect index", err)
		return
	}

	index--

	var new_desc = list.tasks[index].desc
	var new_priority = list.tasks[index].priority

	print("\nNew description (enter for skip): ")
	menu.scanner = bufio.NewScanner(os.Stdin)
	if ok := menu.scanner.Scan(); !ok {
		println("\nIncorrect input")
		return
	}
	menu.input = menu.scanner.Text()
	if menu.input != "" {
		new_desc = menu.input
	}

	println("\nNew priority (enter for skip):")
	println("(l)ow")
	println("(m)edium")
	println("(h)igh")

	print(">> ")
	menu.scanner = bufio.NewScanner(os.Stdin)
	if ok := menu.scanner.Scan(); !ok {
		println("\nIncorrect input")
		return
	}
	menu.input = menu.scanner.Text()

	if menu.input != "" {
		switch menu.input {
		case "l":
			new_priority = low
		case "m":
			new_priority = medium
		case "h":
			new_priority = high
		}
	}

	if index < 0 || index > list.count {
		println("\nIndex cannot be lower than 0 or higher than task count")
		return
	}

	list.tasks[index].Edit(new_desc, new_priority)
	list.WriteToFile()
}

func (menu *Menu) ChangeTaskStatus(list *List) {
	print("\nIndex: ")
	menu.scanner = bufio.NewScanner(os.Stdin)
	if ok := menu.scanner.Scan(); !ok {
		println("\nIncorrect input")
		return
	}
	var index, err = strconv.Atoi(menu.scanner.Text())
	if err != nil {
		println("\nIncorrect index", err)
		return
	}

	if index <= 0 || index > list.count {
		println("\nIndex cannot be lower than 0 or higher than task count")
		return
	}

	list.tasks[index-1].ChangeStatus()
	list.WriteToFile()
}
