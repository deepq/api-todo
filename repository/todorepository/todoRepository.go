package todorepository

import (
	"api-todo/models"
	"fmt"
)

// TodoRepository - In-memory repository
type TodoRepository struct {
	collection []*models.Todo
}

// Init - initializes repository
func (repository *TodoRepository) Init() {
	repository.collection = []*models.Todo{}
}

// Add - adds new item to repository
func (repository *TodoRepository) Add(todo *models.Todo) {
	repository.collection = append(repository.collection, todo)
}

// GetAll - returns all items in repository
func (repository *TodoRepository) GetAll() []*models.Todo {
	return repository.collection
}

// GetByID - finds and return item by its ID
func (repository *TodoRepository) GetByID(id string) (todo *models.Todo, err error) {
	for _, item := range repository.GetAll() {
		if item.ID.String() == id {
			return item, nil
		}
	}

	err = fmt.Errorf("Invalid Id")
	return nil, err
}

// DeleteByID - removes item from collection by its ID
func (repository *TodoRepository) DeleteByID(id string) (err error) {
	for i, item := range repository.GetAll() {
		if item.ID.String() == id {
			//remove
			repository.collection = append(repository.collection[:i], repository.collection[i+1:]...)
			return nil
		}
	}

	err = fmt.Errorf("Invalid Id")
	return err
}

// UpdateByID - updates 'Done' flag by item's ID
func (repository *TodoRepository) UpdateByID(id string, done bool) (err error) {
	for _, item := range repository.GetAll() {
		if item.ID.String() == id {
			item.Done = done
			return nil
		}
	}

	err = fmt.Errorf("Invalid Id")
	return err
}
