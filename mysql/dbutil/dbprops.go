package dbutil

import "fmt"

type Info struct {
	ID   int
	Name string
}

//DbDriver ..
const DbDriver = "mysql"

// User ...
const User = "root"

// password
const Password = "11111111"

// DbName
const DbName = "db_go1"

//Table name ...
const TableName = "person"

//DataSourceName ...
var DataSourceName = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8", User, Password, DbName)
