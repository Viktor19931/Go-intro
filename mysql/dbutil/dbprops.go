package dbutil

import "fmt"

type Info struct {
	ID   int
	Name string
}

// DBMode ...
// const DBMode = "mysql"
const DBMode = "postgres"

//DbDriver ..
// const DbDriver = "mysql"
var DbDriver = "mysql"

// User ...
// const User = "root"
var User = "root"

// password
var Password = "11111111"

// DbName
var DbName = "db_go1"

//Table name ...
const TableName = "person"

//DataSourceName ...
var DataSourceName = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8", User, Password, DbName)
