package app

import (
	"api-todo/lib"
	"api-todo/models"
	"api-todo/repository/todokvrepository"
	"encoding/json"

	"github.com/google/uuid"
)

type App struct {
	//repository todorepository.TodoRepository
	repository todokvrepository.TodoKVRepository
}

func (s *App) Init() {
	s.repository.Init()
	s.repository = todokvrepository.New()
}

func (s *App) GetTodos(req models.ApiRequest, resp *models.ServiceResponse) error {
	resp.Response = s.repository.GetAll()
	return nil
}

func (s *App) GetTodo(req models.ApiRequest, resp *models.ServiceResponse) error {
	//validation
	if len(req.Vars["id"]) == 0 {
		return lib.NotValid
	}

	item, err := s.repository.GetByID(req.Vars["id"])
	if err != nil {
		return err
	}

	resp.Response = item
	return nil
}

func (s *App) DeleteTodo(req models.ApiRequest, resp *models.ServiceResponse) error {
	err := s.repository.DeleteByID(req.Vars["id"])
	if err != nil {
		return err
	}

	return nil
}

func (s *App) CreateTodo(req models.ApiRequest, resp *models.ServiceResponse) error {
	todo := new(models.Todo)
	json.NewDecoder(req.Body).Decode(&todo)
	todo.ID = uuid.New()
	todo.Done = false
	addedTodo := s.repository.Add(*todo)
	resp.Response = addedTodo
	return nil
}

func (s *App) UpdateTodo(req models.ApiRequest, resp *models.ServiceResponse) error {
	var todo models.Todo
	json.NewDecoder(req.Body).Decode(&todo)

	updatedTodo, err := s.repository.UpdateByID(req.Vars["id"], todo)
	if err != nil {
		return err
	}

	resp.Response = updatedTodo
	return nil
}

func (s *App) DeleteCompleted(req models.ApiRequest, resp *models.ServiceResponse) error {
	for _, item := range s.repository.GetAll() {
		if item.Done {
			s.repository.DeleteByID(item.ID.String())
		}
	}
	return nil
}
