package core

import (
	"fmt"
	"os"
	"path"
	"sync"
)

// CreateFile creates a file with the provided filaename in the directory storage/
// If file already exists, its contents are replaced.
func CreateFile(filename string) (string, bool) {
	// get the current relative path
	rootDir, err := os.Getwd()
	if err != nil {
		return "Error accessing directory...", false
	}
	// join current relative file path with name of file
	filename = path.Join(rootDir, "storage", filename)
	// create new file uner the path
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Sprintf("Error to create %s\n", filename), false
	}

	defer file.Close()
	// return name of the file and true if successful
	return filename, true
}


func WriteStringToFile(mutex *sync.Mutex, fileName string,str string) bool {
	return false
}