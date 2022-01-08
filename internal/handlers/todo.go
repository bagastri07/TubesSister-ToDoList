package handlers

import (
	"context"
	"fmt"

	"github.com/bagastri07/TubesSister-ToDoList/internal/repositories"
	todo "github.com/bagastri07/TubesSister-ToDoList/protobuf/go"
)

type TodoServiceServer struct {
	todoRepository *repositories.TodoRepository
}

func NewTodoServiceServer() *TodoServiceServer {
	return &TodoServiceServer{
		todoRepository: repositories.NewTodoRepository(),
	}
}

func (server *TodoServiceServer) Create(ctx context.Context, req *todo.CreateRequest) (*todo.CreateResponse, error) {
	todoRepository := repositories.NewTodoRepository()

	//query
	todoData, err := todoRepository.CreateToDo(req)
	if err != nil {
		return nil, fmt.Errorf("error create todo: %v", err)
	}

	return &todo.CreateResponse{
		Id: todoData.ID,
	}, nil
}

func (server *TodoServiceServer) Read(ctx context.Context, req *todo.ReadRequest) (*todo.ReadResponse, error) {
	todoRepository := repositories.NewTodoRepository()

	todoData, err := todoRepository.ReadTodoByID(req.GetId(), req.GetUser())
	if err != nil {
		return nil, fmt.Errorf("errror read todo: %v", err)
	}

	newToDo := todo.ToDo{
		Id:          todoData.ID,
		Title:       todoData.Title,
		Description: todoData.Description,
		Completed:   todoData.Completed,
		User:        todoData.User,
	}

	return &todo.ReadResponse{
		ToDo: &newToDo,
	}, nil
}

func (server *TodoServiceServer) Update(ctx context.Context, req *todo.UpdateRequest) (*todo.UpdateResponse, error) {
	todoRepository := repositories.NewTodoRepository()

	todoData, err := todoRepository.UpdateToDo(req)
	if err != nil {
		return nil, fmt.Errorf("error update todo: %v", err)
	}

	return &todo.UpdateResponse{
		Updated: todoData.Updated,
	}, nil
}

func (server *TodoServiceServer) Delete(ctx context.Context, req *todo.DeleteRequest) (*todo.DeleteResponse, error) {
	todoRepository := repositories.NewTodoRepository()

	todoData, err := todoRepository.DeleteTodo(req)
	if err != nil {
		return nil, fmt.Errorf("error update todo: %v", err)
	}

	return todoData, nil
}

func (server *TodoServiceServer) ReadAll(ctx context.Context, req *todo.ReadAllRequest) (*todo.ReadAllResponse, error) {
	todoRepository := repositories.NewTodoRepository()

	todoData, err := todoRepository.ReadAll(req)
	if err != nil {
		return nil, fmt.Errorf("error update todo: %v", err)
	}

	return todoData, nil
}

func (server *TodoServiceServer) MarkComplete(ctx context.Context, req *todo.MarkRequest) (*todo.MarkResponse, error) {
	todoRepository := repositories.NewTodoRepository()

	todoData, err := todoRepository.MarkToDo(req)

	if err != nil {
		return nil, fmt.Errorf("error update todo: %v", err)
	}

	return todoData, nil
}
