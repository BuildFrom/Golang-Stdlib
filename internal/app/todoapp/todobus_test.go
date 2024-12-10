package todoapp

import (
	"fmt"
	"testing"
	"time"

	"github.com/BuildFrom/Golang-Stdlib/internal/infrastructure/sqldb"
)

// NewTestTodoRepository is used for integration-style testing with a real database.
func NewTestTodoRepository(db sqldb.Service) TodoRepository {
	return &testTodoRepository{
		db: db,
	}
}

type testTodoRepository struct {
	db sqldb.Service
}

func (r *testTodoRepository) getTodos() ([]Todo, error) {
	// You would normally interact with the db here.
	return []Todo{
		{ID: 1, Title: "Mock Todo 1", Status: "INCOMPLETE"},
		{ID: 2, Title: "Mock Todo 2", Status: "COMPLETE"},
	}, nil
}

func (r *testTodoRepository) getTodoByID(id int) (Todo, error) {
	// Return mock data based on the ID
	if id == 1 {
		return Todo{ID: 1, Title: "Mock Todo 1", Status: "INCOMPLETE"}, nil
	} else if id == 2 {
		return Todo{ID: 2, Title: "Mock Todo 2", Status: "COMPLETE"}, nil
	}

	return Todo{}, fmt.Errorf("Todo not found")
}

func (r *testTodoRepository) createTodo(todo Todo) error {
	// Simulate saving the todo to the database
	return nil
}

func (r *testTodoRepository) updateTodo(todo Todo) error {
	// Simulate updating the todo in the database
	return nil
}

func (r *testTodoRepository) deleteTodo(id int) error {
	// Simulate deleting a todo from the database
	return nil
}

type MockTodoRepository struct {
	GetTodosFunc    func() ([]Todo, error)
	GetTodoByIDFunc func(id int) (Todo, error)
	CreateTodoFunc  func(todo Todo) error
	UpdateTodoFunc  func(todo Todo) error
	DeleteTodoFunc  func(id int) error
}

func (m *MockTodoRepository) getTodos() ([]Todo, error) {
	if m.GetTodosFunc != nil {
		return m.GetTodosFunc()
	}
	return nil, fmt.Errorf("GetTodosFunc not implemented")
}

func (m *MockTodoRepository) getTodoByID(id int) (Todo, error) {
	if m.GetTodoByIDFunc != nil {
		return m.GetTodoByIDFunc(id)
	}
	return Todo{}, fmt.Errorf("GetTodoByIDFunc not implemented")
}

func (m *MockTodoRepository) createTodo(todo Todo) error {
	if m.CreateTodoFunc != nil {
		return m.CreateTodoFunc(todo)
	}
	return fmt.Errorf("CreateTodoFunc not implemented")
}

func (m *MockTodoRepository) updateTodo(todo Todo) error {
	if m.UpdateTodoFunc != nil {
		return m.UpdateTodoFunc(todo)
	}
	return fmt.Errorf("UpdateTodoFunc not implemented")
}

func (m *MockTodoRepository) deleteTodo(id int) error {
	if m.DeleteTodoFunc != nil {
		return m.DeleteTodoFunc(id)
	}
	return fmt.Errorf("DeleteTodoFunc not implemented")
}

func Test_Todo(t *testing.T) {
	t.Parallel()

	// Create a new environment with a mock repository
	mockRepo := &MockTodoRepository{
		GetTodosFunc: func() ([]Todo, error) {
			return []Todo{
				{ID: 1, Title: "Mock Todo 1", Status: "INCOMPLETE"},
				{ID: 2, Title: "Mock Todo 2", Status: "COMPLETE"},
			}, nil
		},
		GetTodoByIDFunc: func(id int) (Todo, error) {
			if id == 1 {
				return Todo{ID: 1, Title: "Mock Todo 1", Status: "INCOMPLETE"}, nil
			} else if id == 2 {
				return Todo{ID: 2, Title: "Mock Todo 2", Status: "COMPLETE"}, nil
			}
			return Todo{}, fmt.Errorf("Todo not found")
		},
		CreateTodoFunc: func(todo Todo) error {
			return nil // Simulate successful creation
		},
		UpdateTodoFunc: func(todo Todo) error {
			return nil // Simulate successful update
		},
		DeleteTodoFunc: func(id int) error {
			return nil // Simulate successful delete
		},
	}

	// Run all test cases
	tests := []struct {
		name string
		run  func(t *testing.T)
	}{
		{"GetTodos", testGetTodos(mockRepo)},
		{"GetTodoByID", testGetTodoByID(mockRepo)},
		{"CreateTodo", testCreateTodo(mockRepo)},
		{"UpdateTodo", testUpdateTodo(mockRepo)},
		{"DeleteTodo", testDeleteTodo(mockRepo)},
	}

	// Execute each test
	for _, tt := range tests {
		t.Run(tt.name, tt.run)
	}
}

// Test GetTodos function
func testGetTodos(repo TodoRepository) func(t *testing.T) {
	return func(t *testing.T) {
		expected := 2
		todos, err := repo.getTodos()
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if len(todos) != expected {
			t.Errorf("Expected %d todos, got %d", expected, len(todos))
		}
	}
}

// Test GetTodoByID function
func testGetTodoByID(repo TodoRepository) func(t *testing.T) {
	return func(t *testing.T) {
		tests := []struct {
			id       int
			expected int
			err      bool
		}{
			{1, 1, false},
			{2, 2, false},
			{999, 0, true},
		}

		for _, tt := range tests {
			t.Run(fmt.Sprintf("Get Todo by ID %d", tt.id), func(t *testing.T) {
				todo, err := repo.getTodoByID(tt.id)
				if tt.err && err == nil {
					t.Errorf("Expected error, got nil")
				}
				if !tt.err && err != nil {
					t.Errorf("Expected no error, got %v", err)
				}
				if todo.ID != tt.expected {
					t.Errorf("Expected ID %d, got %d", tt.expected, todo.ID)
				}
			})
		}
	}
}

// Test CreateTodo function
func testCreateTodo(repo TodoRepository) func(t *testing.T) {
	return func(t *testing.T) {
		todo := Todo{
			Title:     "New Todo",
			Status:    "INCOMPLETE",
			ExpiredAt: parseTime("2024-12-10"),
			CreatedAt: *parseTime("2024-11-22"),
			UpdatedAt: *parseTime("2024-11-22"),
		}
		if err := repo.createTodo(todo); err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	}
}

// Test UpdateTodo function
func testUpdateTodo(repo TodoRepository) func(t *testing.T) {
	return func(t *testing.T) {
		todo := Todo{
			ID:        1,
			Title:     "Updated Todo",
			Status:    "COMPLETE",
			ExpiredAt: parseTime("2024-12-15"),
			CreatedAt: *parseTime("2024-11-01"),
			UpdatedAt: *parseTime("2024-11-01"),
		}
		if err := repo.updateTodo(todo); err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	}
}

// Test DeleteTodo function
func testDeleteTodo(repo TodoRepository) func(t *testing.T) {
	return func(t *testing.T) {
		tests := []int{1, 2}

		for _, id := range tests {
			t.Run(fmt.Sprintf("Delete Todo %d", id), func(t *testing.T) {
				if err := repo.deleteTodo(id); err != nil {
					t.Errorf("Expected no error, got %v", err)
				}
			})
		}
	}
}

// Utility function to parse time from string
func parseTime(timeStr string) *time.Time {
	t, err := time.Parse("2006-01-02", timeStr)
	if err != nil {
		panic(err)
	}
	return &t
}
