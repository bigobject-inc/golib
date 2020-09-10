package dbquery

import "database/sql"

//Database init params
type Database struct {
	ip       string
	port     string
	username string
	password string
	database string

	driverName     string
	dataSourceName string
	DB             *sql.DB
}

// BoQuery init params
type BoQuery struct {
	url string
	ws  string
}
