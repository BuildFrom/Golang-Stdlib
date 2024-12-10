package todoapp

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/BuildFrom/Golang-Stdlib/internal/infrastructure/sqldb"
)

// -----------------------------------------------------------------------------

type TodoRepository interface {
	getTodos() ([]Todo, error)
	getTodoByID(id int) (Todo, error)
	createTodo(todo Todo) error
	updateTodo(todo Todo) error
	deleteTodo(id int) error
}

// -----------------------------------------------------------------------------

type app struct {
	repo TodoRepository
}

func newApp(repo TodoRepository) *app {
	return &app{
		repo: repo,
	}
}

// -----------------------------------------------------------------------------

type store struct {
	db sqldb.Service
}

func newStore(db sqldb.Service) *store {
	return &store{
		db: db,
	}
}

// -----------------------------------------------------------------------------

func (s *store) getTodos() ([]Todo, error) {
	query := `SELECT id, title, status, expires_at AS expired_at, created_at FROM todos`

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Status, &todo.ExpiredAt, &todo.CreatedAt); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func (s *store) getTodoByID(id int) (Todo, error) {
	query := `SELECT id, title, status, expires_at AS expired_at, created_at FROM todos WHERE id = $1`

	var todo Todo
	err := s.db.QueryRow(query, id).Scan(&todo.ID, &todo.Title, &todo.Status, &todo.ExpiredAt, &todo.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return Todo{}, fmt.Errorf("todo with id %d not found", id)
		}
		return Todo{}, err
	}

	return todo, nil
}

func (s *store) createTodo(todo Todo) error {
	query := `INSERT INTO todos (title, status, expires_at) VALUES ($1, $2, $3)`

	_, err := s.db.ExecuteQuery(query, todo.Title, todo.Status, todo.ExpiredAt)
	if err != nil {
		return fmt.Errorf("failed to create todo: %v", err)
	}

	return nil
}

func (s *store) updateTodo(todo Todo) error {
	query := `UPDATE todos SET title = $1, status = $2, expires_at = $3 WHERE id = $4`

	_, err := s.db.ExecuteQuery(query, todo.Title, todo.Status, todo.ExpiredAt, todo.ID)
	if err != nil {
		return fmt.Errorf("failed to update todo with id %d: %v", todo.ID, err)
	}

	return nil
}

func (s *store) deleteTodo(id int) error {
	query := `DELETE FROM todos WHERE id = $1`

	_, err := s.db.ExecuteQuery(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete todo with id %d: %v", id, err)
	}

	return nil
}

// -----------------------------------------------------------------------------

func (a *app) getTodosHandler(w http.ResponseWriter, r *http.Request) {
	todos, err := a.repo.getTodos()
	if err != nil {
		http.Error(w, "Error fetching todos: "+err.Error(), http.StatusInternalServerError)
		return
	}

	encodedData, err := json.Marshal(todos)
	if err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(encodedData)
	if err != nil {
		log.Printf("Error writing response: %v", err)
	}
}

func (a *app) getTodoByIDHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	todo, err := a.repo.getTodoByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	encodedData, err := json.Marshal(todo)
	if err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(encodedData)
	if err != nil {
		log.Printf("Error writing response: %v", err)
	}
}

func (a *app) updateTodoHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var updatedTodo Todo
	err = json.NewDecoder(r.Body).Decode(&updatedTodo)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	updatedTodo.ID = id
	err = a.repo.updateTodo(updatedTodo)
	if err != nil {
		http.Error(w, "Error updating todo: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (a *app) deleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	err = a.repo.deleteTodo(id)
	if err != nil {
		http.Error(w, "Error deleting todo: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (a *app) createTodoHandler(w http.ResponseWriter, r *http.Request) {
	var newTodo Todo
	err := json.NewDecoder(r.Body).Decode(&newTodo)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err = a.repo.createTodo(newTodo)
	if err != nil {
		http.Error(w, "Error creating todo: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
