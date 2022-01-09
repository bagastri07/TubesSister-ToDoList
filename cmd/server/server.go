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
	const grpcGatewayAddress = "0.0.0.0:7001"
	const grpcAddress = "0.0.0.0:7000"

	//grpc gateway
	go func() {
		//mux
		mux := runtime.NewServeMux()
		//register
		protobuf.RegisterToDoServiceHandlerServer(context.Background(), mux, todoServiceServer)

		//http server
		log.Fatalln(http.ListenAndServe(grpcGatewayAddress, mux))
	}()

	lis, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		log.Fatalf("Failed to listen to the port: %v", err)
	}

	//create Server
	server := grpc.NewServer()

	//Register service
	protobuf.RegisterToDoServiceServer(server, todoServiceServer)

	//Run server

	log.Println("gRPC Server is Runing on", grpcAddress)
	if err := server.Serve(lis); err != nil {
		log.Fatal(err.Error())
	}

}
