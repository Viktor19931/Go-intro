package util

import (
	"database/sql"
	"fmt"
)

// Customer
type Customer struct {
	CID         uint    `json:"id"` //Customer ID
	FirstName   string  `json:"first_name"`
	LastName    string  `json: "last_name"`
	CreditScore uint    `json:"credit_score"`
	Salary      float32 `json:"salary"`
	DownPayment float32 `json:"downpayment"`
	HouseID     uint    `json:"house_id"`
}

// House
type House struct {
	HID             uint    `json: "id"`
	Price           float32 `json:"price"`
	MinDownPayment  float32 `json:"min_down"`
	PropertyTax     float32 `json:"prop_tax"`
	MaintenanceCost float32 `json:"m_cost"`
}

// Mortgage
type Mortgage struct {
	Customer
	House
}

// DbDriver
const DbDriver = "mysql"

// User
const User = "root"

// Pass
const Password = "11111111"

// DbName
const DbName = "db_go1"

// customer table
const TableCustomer = "customer"

// house table
const TableHouse = "house"

// DataSourceName ...
var DataSourceName = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8",
	User, Password, DbName)

// DB
var DB *sql.DB

// RowCustomer ...
var RowCustomer *sql.Rows

var NumOfRowCustomer int

// Mrtgs []Mortgage
var Mrtgs []Mortgage

// ApproveIdx ... indices of Mrtgs[] those indicate approved mortgage application
var ApproveIdx []int

// RejectedIdx ... indicates of Mrtgs[] those indicate rejected mortgage applications
var RejectedIdx []int
