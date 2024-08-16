package domain

import "time"

type Todo struct {
	Id          uint      `json:"id" form:"id"`
	Name        string    `json:"name" form:"name"`
	Description string    `json:"description" form:"description"`
	Done        bool      `json:"done" form:"done"`
	CreatedAt   time.Time `json:"created_at" form:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at"`
}

type TodoRepository interface {
	Create(todo *Todo) (*Todo, error)
	GetAll() ([]*Todo, error)
	GetByID(id uint) (*Todo, error)
	Update(todo *Todo) (*Todo, error)
	Delete(id uint) error
}

type TodoService interface {
	AddNewTodo(*Todo) (*Todo, error)
	GetAllTodos() ([]*Todo, error)
	GetTodoById(id uint) (*Todo, error)
	UpdateTodo(todo *Todo) (*Todo, error)
	DeleteTodo(id uint) error
}
