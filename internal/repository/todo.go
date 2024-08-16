package repository

import (
	"errors"
	"go-cleanarch/pkg/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Todo struct {
	gorm.Model

	Name        string `json:"name"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type postgresTodoRepository struct {
	db *gorm.DB
}

func NewPostgresTodoRepository() (domain.TodoRepository, error) {
	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=postgres dbname=todo sslmode=disable"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&Todo{})

	return &postgresTodoRepository{db: db}, nil
}

func (r *postgresTodoRepository) Create(todo *domain.Todo) (*domain.Todo, error) {
	todoModel := Todo{
		Name:        todo.Name,
		Description: todo.Description,
		Done:        todo.Done,
	}

	result := r.db.Create(&todoModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return &domain.Todo{
		Id:          todoModel.ID,
		Name:        todoModel.Name,
		Description: todoModel.Description,
		Done:        todoModel.Done,
		CreatedAt:   todoModel.CreatedAt,
		UpdatedAt:   todoModel.UpdatedAt,
	}, nil
}

func (r *postgresTodoRepository) GetAll() ([]*domain.Todo, error) {
	var todosModel []*Todo
	result := r.db.Find(&todosModel)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, domain.ErrNotFound
	} else if result.Error != nil {
		return nil, result.Error
	}

	var todos []*domain.Todo
	for _, todoModel := range todosModel {
		todos = append(todos, &domain.Todo{
			Id:          todoModel.ID,
			Name:        todoModel.Name,
			Description: todoModel.Description,
			Done:        todoModel.Done,
			CreatedAt:   todoModel.CreatedAt,
			UpdatedAt:   todoModel.UpdatedAt,
		})
	}

	return todos, nil
}

func (r *postgresTodoRepository) GetByID(id uint) (*domain.Todo, error) {
	var todoModel Todo
	result := r.db.First(&todoModel, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, domain.ErrNotFound
	} else if result.Error != nil {
		return nil, result.Error
	}

	todo := &domain.Todo{
		Id:          todoModel.ID,
		Name:        todoModel.Name,
		Description: todoModel.Description,
		Done:        todoModel.Done,
		CreatedAt:   todoModel.CreatedAt,
		UpdatedAt:   todoModel.UpdatedAt,
	}

	return todo, nil
}

func (r *postgresTodoRepository) Update(todo *domain.Todo) (*domain.Todo, error) {
	var todoModel Todo
	todoModel.Model.ID = todo.Id

	result := r.db.
		Model(&todoModel).           // select model (with id)
		Select("*").                 // select all fields to update
		Omit("CreatedAt").           // omit CreatedAt field (avoid to update it)
		Clauses(clause.Returning{}). // return all fields
		Updates(todo)                // update all fields
	if result.Error != nil {
		return nil, result.Error
	}

	return &domain.Todo{
		Id:          todoModel.ID,
		Name:        todoModel.Name,
		Description: todoModel.Description,
		Done:        todoModel.Done,
		CreatedAt:   todoModel.CreatedAt,
		UpdatedAt:   todoModel.UpdatedAt,
	}, nil
}

func (r *postgresTodoRepository) Delete(id uint) error {
	result := r.db.Delete(&Todo{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
