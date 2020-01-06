package todokvrepository

import (
	"api-todo/lib"
	"api-todo/models"
	"sync"
)

type TodoKVRepository struct {
	collection *sync.Map
}

func New() TodoKVRepository {
	return TodoKVRepository{
		collection: &sync.Map{},
	}
}

func (repository *TodoKVRepository) Init() {}

func (repository *TodoKVRepository) Add(todo models.Todo) models.Todo {
	repository.collection.Store(todo.ID.String(), todo)
	return todo
}

func (repository *TodoKVRepository) GetAll() []models.Todo {
	result := []models.Todo{}

	handleRangeItem := func(key, value interface{}) bool {
		todo := value.(models.Todo)
		result = append(result, todo)
		return true
	}

	repository.collection.Range(handleRangeItem)
	return result
}

func (repository *TodoKVRepository) GetByID(id string) (models.Todo, error) {
	value, ok := repository.collection.Load(id)
	if !ok {
		return models.Todo{}, lib.NotFound
	}

	todo, ok := value.(models.Todo)
	if !ok {
		return models.Todo{}, lib.BadItem
	}
	return todo, nil
}

func (repository *TodoKVRepository) DeleteByID(id string) error {
	repository.collection.Delete(id)
	return nil
}

func (repository *TodoKVRepository) UpdateByID(id string, todoToUpdate models.Todo) (models.Todo, error) {
	value, ok := repository.collection.Load(id)
	if !ok {
		return models.Todo{}, lib.NotFound
	}

	todo, ok := value.(models.Todo)
	if !ok {
		return models.Todo{}, lib.BadItem
	}
	todo.Done = todoToUpdate.Done
	todo.Name = todoToUpdate.Name
	repository.collection.Store(todo.ID.String(), todo)
	return todo, nil
}
