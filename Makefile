genmodel:
		gen --sqltype=sqlite3 --connstr "./configs/db/todo.db" --database main --json --gorm
protofile:
		buf generate
start:
		go run main.go