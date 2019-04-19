package main

import (
	"database/sql"
	"fmt"

	dbu "Go-intro/mysql/dbutil"

	_ "github.com/go-sql-driver/mysql"
	// for PostgesSQL
	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

func init() { // for PostgesSQL
	if dbu.DBMode == "postgres" {
		dbu.DbDriver = "postgres"
		dbu.User = "vhordubey"
		dbu.DbName = "go_db1"

		dbu.DataSourceName = fmt.Sprintf("host=localhost port=5432 user=%s "+
			"password=%s dbname=%s sslmode=disable", dbu.User, dbu.Password, dbu.DbName)
	}
}

func main() {
	fmt.Printf("Openng the %s ... \n", dbu.DbName)
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
	// test4Update()
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

	var sqlStmt string
	if dbu.DBMode == "mysql" {
		sqlStmt = fmt.Sprintf("select id, name from %s where id = ?", dbu.TableName)
	} else {
		sqlStmt = fmt.Sprintf("select id, name from %s where id = $1", dbu.TableName)
	}
	err := db.QueryRow(sqlStmt, 3).Scan(&info.ID, &info.Name)
	checkErr(err)

	fmt.Printf("%d\t%s \n", info.ID, info.Name)
}

func test3Insert() {
	fmt.Println("\n ==> test3 insert")

	var sqlStmt string
	if dbu.DBMode == "mysql" {
		sqlStmt = fmt.Sprintf("insert %s set id=?, name=?", dbu.TableName)
	} else {
		sqlStmt = fmt.Sprintf("insert into %s values ($1,$2)", dbu.TableName)
	}

	stmt, err := db.Prepare(sqlStmt)
	checkErr(err)

	var res sql.Result
	res, err = stmt.Exec(7, "Agava")
	checkErr(err)

	// LastInsertId() doesnt work with 'PostgreSQL'
	if dbu.DBMode != "postgres" {
		id, err := res.LastInsertId()
		checkErr(err)

		fmt.Println(id)
	}
}

func test4Update() {
	fmt.Println("\n ==> test4 update")
	var sqlStmt string
	if dbu.DBMode == "mysql" {
		sqlStmt = fmt.Sprintf("update %s set name=? where id=?", dbu.TableName)
	} else {
		sqlStmt = fmt.Sprintf("update %s set name=$1 where id=$2", dbu.TableName)
	}
	stmt, err := db.Prepare(sqlStmt)
	checkErr(err)

	res, err := stmt.Exec("Jack Tree", 2)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
}

func test5Delete() {
	fmt.Println("\n ==> test5 delete")

	var sqlStmt string
	if dbu.DBMode == "mysql" {
		sqlStmt = fmt.Sprintf("delete from %s where id=?", dbu.TableName)
	} else {
		sqlStmt = fmt.Sprintf("delete from %s where id=$1", dbu.TableName)
	}
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
		panic(err.Error())
	}
}
