package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

func AddUsersToMysqlDB(count int, db *sql.DB) {

	stmtIns, err := db.Prepare("INSERT INTO users (age, name) VALUES( ?, ? )") // ? = placeholder
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()

	for i := 0; i < count; i++ {
		_, err := stmtIns.Exec(i, "bar"+strconv.Itoa(i))
		if err != nil {
			panic(err.Error())
		}
	}
}
