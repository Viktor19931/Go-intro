package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"runtime"
	"time"

	utl "Go-intro/project/util"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	rand.Seed(time.Now().UnixNano())

	maxCPU := runtime.NumCPU()

	utl.CPUUsed = 8
	runtime.GOMAXPROCS(utl.CPUUsed)

	fmt.Printf("\n======================================\n")
	fmt.Printf("= Number of CPUs (Total=%d - Used=%d)", maxCPU, utl.CPUUsed)
	fmt.Printf("\n======================================\n\n")
}

func main() {

	fmt.Printf("Opening the %s ...\n\n", utl.DbName)
	var err error
	utl.DB, err = sql.Open(utl.DbDriver, utl.DataSourceName)

	defer utl.DB.Close()

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Success !!")
	}

	utl.MatchustomerHouses()
	utl.LogMrtgApps()
}
