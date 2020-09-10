package dbquery

import (
	"database/sql"
	"reflect"
)

// Dbquery service
type Dbquery interface {
	ConnectDB() error
	CloseDB()
	Query(sql string, tblstruc interface{}) ([]reflect.Value, error)
	GetData(row interface{}, column string) interface{}
	GetDB() *sql.DB
}
