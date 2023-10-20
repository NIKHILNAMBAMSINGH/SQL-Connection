package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const DBName = "learning"
const DbUser = "root"
const DbPassword = "ril29856884N"

type Data struct {
	id   int
	name string
}

func checkError(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}

func main() {

	connectionString := fmt.Sprintf("%v:%v@tcp(localhost:3306)/%v", DbUser, DbPassword, DBName)
	db, err := sql.Open("mysql", connectionString)
	checkError(err)
	defer db.Close()

	result, err := db.Exec("insert into data values(10,'oditya')")
	checkError(err)
	lastInsertedId, err := result.LastInsertId()
	fmt.Println("lastInsertedID", lastInsertedId)
	checkError(err)
	rowsAffected, err := result.RowsAffected()
	fmt.Println("rowsAffected", rowsAffected)
	checkError(err)

	rows, err := db.Query("SELECT * from data")
	checkError(err)

	for rows.Next() {
		var data Data
		err := rows.Scan(&data.id, &data.name)
		checkError(err)
		fmt.Println(data)
	}
}
