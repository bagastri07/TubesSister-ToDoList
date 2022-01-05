package repositories

import (
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

func (r *TodoRepository) ReadTodoByID(todoId int32) (*model.Todo, error) {
	var todoData model.Todo

	err := r.dbClient.First(&todoData, todoId)

	if err.Error != nil {
		return nil, err.Error
	}

	return &todoData, nil
}

func (r *TodoRepository) CreateToDo(req *todo.CreateRequest) (*model.Todo, error) {
	todoData := model.Todo{
		Title:       req.ToDo.GetTitle(),
		Description: req.ToDo.GetDescription(),
		Completed:   0,
	}

	if err := r.dbClient.Create(&todoData).Error; err != nil {
		return nil, err
	}

	return &todoData, nil

}

func (r *TodoRepository) UpdateToDo(req *todo.UpdateRequest) (*todo.UpdateResponse, error) {
	var todoData model.Todo

	if err := r.dbClient.Where("ID = ?", req.GetId()).First(&todoData).Error; err != nil {
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
	var todos []model.Todo

	if err := r.dbClient.Find(&todos).Error; err != nil {
		return nil, err
	}

	todosList := []*todo.ToDo{}
	for i := 0; i < len(todos); i++ {
		td := todo.ToDo{
			Id:          todos[i].ID,
			Title:       todos[i].Title,
			Description: todos[i].Description,
			Completed:   todos[i].Completed,
		}
		todosList = append(todosList, &td)
	}

	return &todo.ReadAllResponse{
		ToDos: todosList,
	}, nil
}

func (r *TodoRepository) DeleteTodo(req *todo.DeleteRequest) (*todo.DeleteResponse, error) {
	var todoData model.Todo

	result := r.dbClient.Delete(&todoData, req.GetId())

	if err := result.Error; err != nil {
		return nil, err
	}

	return &todo.DeleteResponse{
		Deleted: int32(result.RowsAffected),
	}, nil
}

func (r *TodoRepository) MarkToDo(req *todo.MarkRequest) (*todo.MarkResponse, error) {
	var todoData model.Todo

	if err := r.dbClient.Where("ID = ?", req.Id).First(&todoData).Error; err != nil {
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
