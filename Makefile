genmodel:
		gen --sqltype=sqlite3 --connstr "./configs/db/todo.db" --database main --json --gorm
protofile:
		buf generate
start-server:
		go run cmd/server/server.go
start-client:
		go run cmd/client/client.go