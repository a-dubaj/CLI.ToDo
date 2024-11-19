package main

func main() {
	todos := Todos{}
	todos.add("Buy milk") // TODO: for test only
	todos.add("Buy tea")  // TODO: for test only
	todos.toggle(0)
	todos.print()
}
