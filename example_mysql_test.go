package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tborisova/clean_like_gopher"
	"testing"
)

var mysqlStartOptions = map[string]string{"username": "root", "dbName": "golangtest"}

func TestAdd25UsersToDb(t *testing.T) {
	cleaner, _ := clean_like_gopher.NewCleaningGopher("mysql", mysqlStartOptions)
	db, _ := sql.Open("mysql", "root:@/golangtest")
	defer db.Close()

	AddUsersToMysqlDB(25, db)

	var count int
	row := db.QueryRow("SELECT count(id) as count FROM users")
	_ = row.Scan(&count)

	if count != 25 {
		t.Errorf("Expected 25 users to be added!")
	}

	cleaner.Clean(nil)
	cleaner.Close()
}

func TestAdd5UsersToDb(t *testing.T) {
	cleaner, _ := clean_like_gopher.NewCleaningGopher("mysql", mysqlStartOptions)
	db, _ := sql.Open("mysql", "root:@/golangtest")
	defer db.Close()

	AddUsersToMysqlDB(20, db)

	var count int
	row := db.QueryRow("SELECT count(id) as count FROM users")
	_ = row.Scan(&count)

	if count != 20 {
		t.Errorf("Expected 20 users to be added!")
	}

	cleaner.Clean(nil)
	cleaner.Close()
}
