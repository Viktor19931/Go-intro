package util

import "sync"

// CPUUsed
var CPUUsed int

// WaitG
var WaitG sync.WaitGroup

// Mu..
var Mu sync.Mutex

// CheckErr..
func CheckErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}