package main

import (
	"database/sql"
	"fmt"

	dbu "Go-intro/mysql/dbutil"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func main() {
	fmt.Printf("Openng the %s ...", dbu.DbName)
	db, err = sql.Open(dbu.DbDriver, dbu.DataSourceName)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Success !")
	}

	defer db.Close()

	test1Select()
	test2Select()
	// test3Insert()
	//test4Update()
	test5Delete()
}

func test1Select() {
	fmt.Println("\n ==> test1 select")

	result, err := db.Query("select id, name from " + dbu.TableName)
	checkErr(err)

	for result.Next() {
		var info dbu.Info

		err = result.Scan(&info.ID, &info.Name)
		checkErr(err)

		fmt.Printf("%d\t%s \n", info.ID, info.Name)
	}
}

func test2Select() {
	fmt.Println("\n ==> test2 select")
	var info dbu.Info

	var sqlStmt = fmt.Sprintf("select id, name from %s where id = ?", dbu.TableName)
	err := db.QueryRow(sqlStmt, 3).Scan(&info.ID, &info.Name)
	checkErr(err)

	fmt.Printf("%d\t%s \n", info.ID, info.Name)
}

func test3Insert() {
	fmt.Println("\n ==> test3 insert")

	var sqlStmt = fmt.Sprintf("insert %s set id=?, name=?", dbu.TableName)
	stmt, err := db.Prepare(sqlStmt)
	checkErr(err)

	res, err := stmt.Exec(6, "Anna")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
}

func test4Update() {
	fmt.Println("\n ==> test4 update")

	var sqlStmt = fmt.Sprintf("update %s set name=? where id=?", dbu.TableName)
	stmt, err := db.Prepare(sqlStmt)
	checkErr(err)

	res, err := stmt.Exec("Jack", 2)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
}

func test5Delete() {
	fmt.Println("\n ==> test5 delete")

	var sqlStmt = fmt.Sprintf("delete from %s where id=?", dbu.TableName)
	stmt, err := db.Prepare(sqlStmt)
	checkErr(err)

	res, err := stmt.Exec(5)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error)
	}
}
