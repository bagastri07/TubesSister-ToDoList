package model

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
	uuid "github.com/satori/go.uuid"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
	_ = uuid.UUID{}
)

/*
DB Table Details
-------------------------------------


CREATE TABLE `todo` (`ID` integer,`Title` text,`Description` text,`Completed` integer DEFAULT 0,PRIMARY KEY (`ID`))

JSON Sample
-------------------------------------
{    "id": 85,    "title": "MNHuEmtmIVRmfhAKoGpavutsM",    "description": "DoBXfphMKYZcFIKvbsZnsqlnF",    "completed": 44}



*/

// Todo struct is a row record of the todo table in the main database
type Todo struct {
	//[ 0] ID                                             integer              null: false  primary: true   isArray: false  auto: false  col: integer         len: -1      default: []
	ID int32 `gorm:"primary_key;column:ID;type:integer;" json:"id"`
	//[ 1] Title                                          text                 null: true   primary: false  isArray: false  auto: false  col: text            len: -1      default: []
	Title string `gorm:"column:Title;type:text;" json:"title"`
	//[ 2] Description                                    text                 null: true   primary: false  isArray: false  auto: false  col: text            len: -1      default: []
	Description string `gorm:"column:Description;type:text;" json:"description"`
	//[ 3] Completed                                      integer              null: true   primary: false  isArray: false  auto: false  col: integer         len: -1      default: [0]
	Completed int32 `gorm:"column:Completed;type:integer;default:0;" json:"completed"`
}

var todoTableInfo = &TableInfo{
	Name: "todo",
	Columns: []*ColumnInfo{

		{
			Index:              0,
			Name:               "ID",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "integer",
			DatabaseTypePretty: "integer",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "integer",
			ColumnLength:       -1,
			GoFieldName:        "ID",
			GoFieldType:        "int32",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		{
			Index:              1,
			Name:               "Title",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       -1,
			GoFieldName:        "Title",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "title",
			ProtobufFieldName:  "title",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		{
			Index:              2,
			Name:               "Description",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       -1,
			GoFieldName:        "Description",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "description",
			ProtobufFieldName:  "description",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		{
			Index:              3,
			Name:               "Completed",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "integer",
			DatabaseTypePretty: "integer",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "integer",
			ColumnLength:       -1,
			GoFieldName:        "Completed",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "completed",
			ProtobufFieldName:  "completed",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},
	},
}

// // TableName sets the insert table name for this struct type
// func (t *Todo) TableName() string {
// 	return "todo"
// }

// // BeforeSave invoked before saving, return an error if field is not populated.
// func (t *Todo) BeforeSave() error {
// 	return nil
// }

// // Prepare invoked before saving, can be used to populate fields etc.
// func (t *Todo) Prepare() {
// }

// // Validate invoked before performing action, return an error if field is not populated.
// func (t *Todo) Validate(action Action) error {
// 	return nil
// }

// // TableInfo return table meta data
// func (t *Todo) TableInfo() *TableInfo {
// 	return todoTableInfo
// }
