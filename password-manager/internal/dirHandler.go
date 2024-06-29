package internal

import (
	"os"
)

var (
	dataPath = "data"
)

func VerifyDataFile() bool {
	dirs, _ := os.ReadDir(".")
	var found = false
	for _, dir := range dirs {
		if dir.IsDir() && dataPath == dir.Name() {
			found = true
		}
	}
	return found
}

func CreateDataFile() {
	err := os.Mkdir(dataPath, 0666)
	if err != nil {
		os.Exit(1)
	}
}

func InitializeSubFile(filename string) *os.File {
	file, err := os.Open(dataPath + "/" + filename)
	if err != nil {
		file1, err1 := os.Create(dataPath + "/" + filename)
		if err1 != nil {
			os.Exit(1)
		}
		file = file1
	}
	return file
}
