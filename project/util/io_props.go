package util

import (
	"os"
	"strings"
)

// MrtgType
type MrtgType []string

// FileName
const FileName = "mrtg.txt"

var fileHandler *os.File

// OpenFile
func OpenFile(fileName string) {
	var err error
	fileHandler, err =
		os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0664)

	CheckErr(err)
}

// CloseFile
func CloseFile() {
	err := fileHandler.Close()
	CheckErr(err)
}

// WriteToFile
func WriteToFile(mrtgSummary MrtgType) {
	fileHandler.Write([]byte(mrtgSummary.toString()))
	fileHandler.Write([]byte("\n"))
}

func (t MrtgType) toString() string {
	return strings.Join([]string(t), ",")
}