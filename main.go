package main

import "fmt"

func main() {
	todos := Todos{}
	todos.add("Buy milk") // TODO: for test only
	todos.add("Buy tea")  // TODO: for test only

	fmt.Printf("%+v\n\n", todos)
	todos.delete(0)
	fmt.Printf("%+v", todos)
}
