package db

import (
	"github.com/bagastri07/TubesSister-ToDoList/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GetDBConnection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("./configs/db/todo.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.Todo{})

	return db
}
