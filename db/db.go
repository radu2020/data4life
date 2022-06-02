package db

import (
	"../utils"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

// Conn returns a DB connection after creating the table
func OpenDbConn() (*sql.DB, error){
	// Open DB connection
	conn, err := sql.Open("sqlite3", "./data4life.db")
	utils.Must(err)

	// Create DB Table
	statement, _ := conn.Prepare(
		"CREATE TABLE IF NOT EXISTS tokens (" +
			"id INTEGER PRIMARY KEY AUTOINCREMENT," +
			"token TEXT," +
			"repeated INTEGER," + // sqlite boolean is an int (0 = false, 1 = true)
			"frequency INTEGER" +
			")")
	_, err = statement.Exec()
	utils.Must(err)

	return conn, err
}
