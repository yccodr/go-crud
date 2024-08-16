package service

import (
	"errors"
	"go-cleanarch/pkg/domain"
	"log/slog"

	"gorm.io/gorm"
)

type TodoService struct {
	todoRepository domain.TodoRepository
}

func NewTodoService(todoRepository domain.TodoRepository) *TodoService {
	return &TodoService{todoRepository: todoRepository}
}

func (s *TodoService) AddNewTodo(todo *domain.Todo) (*domain.Todo, error) {
	todo, err := s.todoRepository.Create(todo)

	if err != nil {
		slog.Error("[Service][Todo] AddNewTodo", "err", err)
		return nil, err
	}

	return todo, nil
}

func (s *TodoService) GetAllTodos() ([]*domain.Todo, error) {
	todos, err := s.todoRepository.GetAll()
	if err != nil {
		slog.Error("[Service][Todo] GetAllTodos", "err", err)
		return nil, err
	}

	return todos, nil
}

func (s *TodoService) GetTodoById(id uint) (*domain.Todo, error) {
	todo, err := s.todoRepository.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}

		slog.Error("[Service][Todo] GetTodoById", "err", err)
		return nil, err
	}

	return todo, nil
}

func (s *TodoService) UpdateTodo(todo *domain.Todo) (*domain.Todo, error) {
	todo, err := s.todoRepository.Update(todo)
	if err != nil {
		slog.Error("[Service][Todo] UpdateTodo", "err", err)
		return nil, err
	}

	return todo, nil
}

func (s *TodoService) DeleteTodo(id uint) error {
	err := s.todoRepository.Delete(id)
	if err != nil {
		slog.Error("[Service][Todo] DeleteTodo", "err", err)
		return err
	}

	return nil
}
