package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/bagastri07/TubesSister-ToDoList/internal/handlers"
	protobuf "github.com/bagastri07/TubesSister-ToDoList/protobuf/go"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func main() {

	todoServiceServer := handlers.NewTodoServiceServer()

	go func() {
		//mux
		mux := runtime.NewServeMux()
		//register
		protobuf.RegisterToDoServiceHandlerServer(context.Background(), mux, todoServiceServer)

		//http server
		log.Fatalln(http.ListenAndServe("localhost:7001", mux))
	}()

	lis, err := net.Listen("tcp", "localhost:7000")
	if err != nil {
		log.Fatalf("Failed to listen to the port: %v", err)
	}

	//create Server
	server := grpc.NewServer()

	//Register service
	protobuf.RegisterToDoServiceServer(server, todoServiceServer)

	//Run server
	if err := server.Serve(lis); err != nil {
		log.Fatal(err.Error())
	}
}
