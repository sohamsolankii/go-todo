package main

func main() {
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")

	storage.Load(&todos)

	cmdFlags := NewCmdFlag()
	cmdFlags.execute(&todos)

	storage.Save(todos)

}
