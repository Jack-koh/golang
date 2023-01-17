package model

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type sqlHandler struct {
	db *sql.DB
}

func (s *sqlHandler) GetTodos(sessionId string) []*Todo {
	todos := []*Todo{}
	rows, err := s.db.Query("SELECT * FROM todos WHERE sessionId=?", sessionId)
	defer rows.Close()
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var todo Todo
		rows.Scan(&todo.ID, &todo.Name, &todo.Completed, &todo.CreatedAt)
		todos = append(todos, &todo)
	}
	return todos
}

func (s *sqlHandler) AddTodo(name string, sessionId string) *Todo {
	statement, err := s.db.Prepare("INSERT INTO todos (sessionId, name, completed,createdAt) VALUES (?, ?, ?, DATE_FORMAT(now(), '%Y-%m-%d %H:%i:%s'))")
	if err != nil {
		panic(err)
	}
	result, err := statement.Exec(sessionId, name, false)
	if err != nil {
		panic(err)
	}
	id, _ := result.LastInsertId()
	var todo Todo
	todo.ID = int(id)
	todo.Name = name
	todo.Completed = false
	todo.CreatedAt = time.Now()
	return &todo
}

func (s *sqlHandler) RemoveTodo(id int) bool {
	statement, err := s.db.Prepare("DELETE FROM todos WHERE id = ?")
	if err != nil {
		panic(err)
	}
	result, err := statement.Exec(id)
	if err != nil {
		panic(err)
	}
	cnt, _ := result.RowsAffected()
	return cnt > 0
}

func (s *sqlHandler) CompleteTodo(id int, complete bool) bool {
	statement, err := s.db.Prepare("UPDATE todos SET completed = ? WHERE id = ?")
	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}
	result, err := statement.Exec(complete, id)
	cnt, _ := result.RowsAffected()
	return cnt > 0
}

func (s *sqlHandler) Close() {
	s.db.Close()
}
func newSqlHandler() DBHandler {
	database, err := sql.Open("mysql", "root:qwe123qwe123!@/todo")
	if err != nil {
		panic(err)
	}
	database.Query(`DROP TABLE IF EXISTS todos;`)
	database.Query(`DROP INDEX sessionIdIndexOnTodos ON Todos;`)

	statement, _ := database.Prepare(
		`CREATE TABLE IF NOT EXISTS todos (
					id INT AUTO_INCREMENT PRIMARY KEY,
					sessionId VARCHAR(255),
					name VARCHAR(255),
					completed BOOLEAN,
					createdAt DATETIME
				);`)
	statement.Exec()
	database.Query(`CREATE INDEX sessionIdIndexOnTodos ON todos (sessionId ASC);`)

	return &sqlHandler{db: database}
}
