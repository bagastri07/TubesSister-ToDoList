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

func menuTodoList(user string) {
	fmt.Println("You're Log In as", user)

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
	fmt.Print("Tekan Enter Untuk Kembali Ke Beranda ...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func inputWithSpace() string {
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	name = strings.TrimSuffix(name, "\n")

	return name
}

func showError(err error) {
	fmt.Printf("\n%s \n\n", err.Error())
	pressEnterMenu()
}

func main() {

	//create Connection to gRPC server
	conn, err := grpc.Dial(todoServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to service: %v", err)
		return
	}
	defer conn.Close()

	log.Println("Client has connected to the server on :", todoServiceAddress)

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
		menuTodoList(user)
		fmt.Scanln(&input)
		fmt.Println()
		switch input {
		case 1:
			fmt.Println("Membuat Todo Baru")
			fmt.Println()

			fmt.Print("Masukan Title     : ")
			title := inputWithSpace()

			fmt.Print("Masukan Deskripsi : ")
			desc := inputWithSpace()

			res, err := todoServiceClient.Create(ctx, &todo.CreateRequest{
				ToDo: &todo.ToDo{
					Title:       title,
					Description: desc,
					Completed:   0,
					User:        user,
				},
			})

			if err != nil {
				showError(err)
				break
			}

			fmt.Println("Todo Berhasil dibuat, id: ", res.GetId())
			pressEnterMenu()
		case 2:
			fmt.Println("Tampilkan semua Todo List")
			fmt.Println()

			res, err := todoServiceClient.ReadAll(ctx, &todo.ReadAllRequest{
				User: user,
			})

			if err != nil {
				showError(err)
				break
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

			fmt.Print("\nMasukan ID Todo anda: ")
			var id int32
			fmt.Scanln(&id)

			res, err := todoServiceClient.Read(ctx, &todo.ReadRequest{
				Id:   id,
				User: user,
			})

			if err != nil {
				showError(err)
				break

			}

			fmt.Println("ID             :", res.GetToDo().Id)
			fmt.Println("Title          :", res.GetToDo().Title)
			fmt.Println("Description    :", res.GetToDo().Description)
			if res.GetToDo().Completed == 0 {
				fmt.Println("Completed      : Undone")
			} else {
				fmt.Println("Completed      : Done")
			}
			fmt.Println()

			pressEnterMenu()
		case 4:
			fmt.Println("Mengubah Todo List Berdasarkan ID")

			var id int32
			fmt.Print("Masukan ID Todo : ")
			fmt.Scanln(&id)
			fmt.Println("Masukan data baru")
			fmt.Print("Title          : ")
			title := inputWithSpace()
			fmt.Print("Description    : ")
			description := inputWithSpace()

			res, err := todoServiceClient.Update(ctx, &todo.UpdateRequest{
				Id:   id,
				User: user,
				ToDo: &todo.ToDo{
					Title:       title,
					Description: description,
				},
			})

			if err != nil {
				showError(err)
				break
			}

			fmt.Println()
			fmt.Println("Rowaffected:", res.GetUpdated())
			fmt.Println("Todo List dengan id", id, "berhasil diupdate!")
			fmt.Println()
			pressEnterMenu()
		case 5:
			fmt.Println("Menandai Todo List (Mark as Done) Berdasarkan ID")
			var id int32
			fmt.Print("Masukan ID Todo anda: ")
			fmt.Scanln(&id)
			res, err := todoServiceClient.MarkComplete(ctx, &todo.MarkRequest{
				Id:   id,
				User: user,
			})

			if err != nil {
				showError(err)
				break
			}

			fmt.Println("marsds:", res.GetMarkedId())
			fmt.Println("Todo List dengan id", id, "berhasil ditandai!")
			fmt.Println()
			pressEnterMenu()

		case 6:
			fmt.Println("Menghapus Todo List Berdasarkan ID")
			var id int32
			fmt.Print("Masukan ID Todo anda: ")
			fmt.Scanln(&id)

			res, err := todoServiceClient.Delete(ctx, &todo.DeleteRequest{
				Id:   id,
				User: user,
			})

			if err != nil {
				showError(err)
				break
			}

			fmt.Println("rawAffected:", res.GetDeleted())
			fmt.Println("Todo List dengan id", id, "berhasil dihapus!")
			fmt.Println()
			pressEnterMenu()
		default:
			if input == 7 {
				fmt.Println("Terimakasih, sampai jumpa", user)
			} else {
				fmt.Printf("Input salah, ulangi!!! \n\n")
			}
		}
	}
}
