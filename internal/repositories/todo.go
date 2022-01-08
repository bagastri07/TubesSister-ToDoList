package repositories

import (
	"errors"
	"fmt"

	"github.com/bagastri07/TubesSister-ToDoList/configs/db"
	"github.com/bagastri07/TubesSister-ToDoList/model"
	todo "github.com/bagastri07/TubesSister-ToDoList/protobuf/go"
	"gorm.io/gorm"
)

type TodoRepository struct {
	dbClient *gorm.DB
}

func NewTodoRepository() *TodoRepository {
	dbClient := db.GetDBConnection()
	return &TodoRepository{
		dbClient: dbClient,
	}
}

func (r *TodoRepository) ReadTodoByID(todoId int32, user string) (*model.Todos, error) {
	var todoData model.Todos

	err := r.dbClient.First(&todoData, todoId)

	if err.Error != nil {
		return nil, err.Error
	}

	if todoData.User != user {
		return nil, errors.New("Access Denied! - Wrong User")
	}

	return &todoData, nil
}

func (r *TodoRepository) CreateToDo(req *todo.CreateRequest) (*model.Todos, error) {
	todoData := model.Todos{
		Title:       req.ToDo.GetTitle(),
		Description: req.ToDo.GetDescription(),
		Completed:   0,
		User:        req.ToDo.GetUser(),
	}
	fmt.Println(req.ToDo.GetUser())

	if err := r.dbClient.Create(&todoData).Error; err != nil {
		return nil, err
	}

	return &todoData, nil

}

func (r *TodoRepository) UpdateToDo(req *todo.UpdateRequest) (*todo.UpdateResponse, error) {
	var todoData model.Todos

	if err := r.dbClient.Where("ID = ? AND User = ?", req.GetId(), req.GetUser()).First(&todoData).Error; err != nil {
		return nil, err
	}

	todoData.Title = req.ToDo.Title
	todoData.Description = req.ToDo.Description
	todoData.Completed = req.ToDo.Completed

	result := r.dbClient.Save(&todoData)

	if err := result.Error; err != nil {
		return nil, err
	}

	return &todo.UpdateResponse{
		Updated: int32(result.RowsAffected),
	}, nil
}

func (r *TodoRepository) ReadAll(req *todo.ReadAllRequest) (*todo.ReadAllResponse, error) {
	var todos []model.Todos

	if err := r.dbClient.Where("User = ?", req.GetUser()).Find(&todos).Error; err != nil {
		return nil, err
	}

	todosList := []*todo.ToDo{}
	for i := 0; i < len(todos); i++ {
		td := todo.ToDo{
			Id:          todos[i].ID,
			Title:       todos[i].Title,
			Description: todos[i].Description,
			Completed:   todos[i].Completed,
			User:        todos[i].User,
		}
		todosList = append(todosList, &td)
	}

	return &todo.ReadAllResponse{
		ToDos: todosList,
	}, nil
}

func (r *TodoRepository) DeleteTodo(req *todo.DeleteRequest) (*todo.DeleteResponse, error) {
	var todoData model.Todos

	result := r.dbClient.Where("ID = ? AND User = ?", req.GetId(), req.GetUser()).Delete(&todoData)

	if err := result.Error; err != nil {
		return nil, err
	}

	return &todo.DeleteResponse{
		Deleted: int32(result.RowsAffected),
	}, nil
}

func (r *TodoRepository) MarkToDo(req *todo.MarkRequest) (*todo.MarkResponse, error) {
	var todoData model.Todos

	if err := r.dbClient.Where("ID = ? AND User = ?", req.GetId(), req.GetUser()).First(&todoData).Error; err != nil {
		return nil, err
	}

	todoData.Completed = 1

	result := r.dbClient.Save(&todoData)

	if err := result.Error; err != nil {
		return nil, err
	}

	return &todo.MarkResponse{
		MarkedId: todoData.ID,
	}, nil
}
