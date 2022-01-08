package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	todo "github.com/bagastri07/TubesSister-ToDoList/protobuf/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const todoServiceAddress = "localhost:7000"

func menuTodoList() {
	fmt.Println("1. Buat Todo Baru")
	fmt.Println("2. Tampilkan Semua Todo List")
	fmt.Println("3. Tampilakn Todo List Berdasarkan ID")
	fmt.Println("4. Ubah Todo")
	fmt.Println("5. Tandai Todo (Mark as Done)")
	fmt.Println("6. Hapus Todo")
	fmt.Println("7. Keluar")
	fmt.Println("==========================")
	fmt.Print("Input: ")
}

func pressEnterMenu() {
	fmt.Print("Tekan Enter Untuk Kembali Ke Beranda")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func inputWithString() string {
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	name = strings.TrimSuffix(name, "\n")

	return name
}

func main() {

	//create Connection to gRPC server
	conn, err := grpc.Dial(todoServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to service: %v", err)
		return
	}
	defer conn.Close()

	//create New Service
	todoServiceClient := todo.NewToDoServiceClient(conn)

	//Create connection timeout
	ctx := context.Background()

	//deklarasi variabel user
	var user string
	var input int32

	fmt.Print("Masukan Username anda : ")
	fmt.Scanln(&user)

	input = 0

	for input != 7 {
		menuTodoList()
		fmt.Scanln(&input)
		switch input {
		case 1:
			fmt.Println("Membuat Todo Baru \n")

			fmt.Print("Masukan Title     : ")
			title := inputWithString()

			fmt.Print("Masukan Deskripsi : ")
			desc := inputWithString()

			res, err := todoServiceClient.Create(ctx, &todo.CreateRequest{
				ToDo: &todo.ToDo{
					Title:       title,
					Description: desc,
					Completed:   0,
					User:        user,
				},
			})

			if err != nil {
				panic(err)
			}

			fmt.Println("Todo Berhasil dibuat, id: ", res.GetId())
			pressEnterMenu()
		case 2:
			fmt.Println("Tampilkan semua Todo List \n")

			res, err := todoServiceClient.ReadAll(ctx, &todo.ReadAllRequest{
				User: user,
			})

			if err != nil {
				panic(err)
			}

			if len(res.GetToDos()) == 0 {
				fmt.Println("Tidak ada Data Todo Untuk User: ", user)
			}

			for i := 0; i < len(res.GetToDos()); i++ {
				fmt.Println("ID             :", res.GetToDos()[i].Id)
				fmt.Println("Title          :", res.GetToDos()[i].Title)
				fmt.Println("Description    :", res.GetToDos()[i].Description)
				if res.GetToDos()[i].Completed == 0 {
					fmt.Println("Completed      : Undone")
				} else {
					fmt.Println("Completed      : Done")
				}
				fmt.Println()
			}
			pressEnterMenu()
		case 3:
			fmt.Println("Menampilkan Todo List Berdasarkan ID")
			fmt.Println()

			fmt.Println("Masukan ID Todo anda: ")
			var id int32
			fmt.Scanln(&id)

			res, err := todoServiceClient.Read(ctx, &todo.ReadRequest{
				Id:   id,
				User: user,
			})

			if err != nil {
				panic(err)
			}

			fmt.Println(res.GetToDo().Id)
			fmt.Println(res.GetToDo().Title)
			fmt.Println(res.GetToDo().Description)
			fmt.Println(res.GetToDo().Completed)

			pressEnterMenu()
		case 4:
			fmt.Println("Mengubah Todo List Berdasarkan ID")

			var id int32
			fmt.Println("Masukan Judul Baru Todo anda: ")
			title := inputWithString()
			fmt.Println("Masukan Deskripsi Baru Todo anda: ")
			description := inputWithString()

			res, err := todoServiceClient.Update(ctx, &todo.UpdateRequest{
				Id:   id,
				User: user,
				ToDo: &todo.ToDo{
					Title: title,
					Description: description,
				},
			})

			if err != nil {
				panic(err)
			}

			fmt.Println("Rowaffected:", res.GetUpdated())
			fmt.Println("Todo List dengan id", id, "berhasil diupdate!")
			pressEnterMenu()
		case 5:
			fmt.Println("Menandai Todo List (Mark as Done) Berdasarkan ID")
			var id int32
			fmt.Println("Masukan ID Todo anda: ")
			fmt.Scanln(&id)
			res, err := todoServiceClient.MarkComplete(ctx, &todo.MarkRequest{
				Id:   id,
				User: user,
			})

			if err != nil {
				panic(err)
			}

			fmt.Println("marsds:", res.GetMarkedId())
			fmt.Println("Todo List dengan id", id, "berhasil ditandai!")
			pressEnterMenu()

		case 6:
			fmt.Println("Menghapus Todo List Berdasarkan ID")
			var id int32
			fmt.Println("Masukan ID Todo anda: ")
			fmt.Scanln(&id)

			res, err := todoServiceClient.Delete(ctx, &todo.DeleteRequest{
				Id:   id,
				User: user,
			})

			if err != nil {
				panic(err)
			}

			fmt.Println("rawAffected:", res.GetDeleted())
			fmt.Println("Todo List dengan id", id, "berhasil dihapus!")
			pressEnterMenu()
		}
	}
}
