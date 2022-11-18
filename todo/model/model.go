package model

import "time"

type Todo struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

type DBHandler interface {
	GetTodos() []*Todo
	AddTodo(name string) *Todo
	RemoveTodo(id int) bool
	CompleteTodo(id int, complete bool) bool
	Close()
}

func NewDBHandler() DBHandler {
	return newSqlHandler()
}

/*
func init() {
	//handler = newMemoryHandler()
	handler = newSqlHandler()
}

func GetTodos() []*Todo {
	return handler.GetTodos()
}

func AddTodo(name string) *Todo {
	return handler.AddTodo(name)
}

func RemoveTodo(id int) bool {
	return handler.RemoveTodo(id)
}

func CompleteTodo(id int, complete bool) bool {
	return handler.CompleteTodo(id, complete)
}
*/
