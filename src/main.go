package main

func main() {
	var list = List{count: 0}
	var menu = Menu{input: ""}

	list.Load() // Load from the save

	for menu.input != "q" {
		menu.Draw(list)
		menu.HandleInput(&list)
	}
}
