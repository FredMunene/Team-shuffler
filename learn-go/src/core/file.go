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

func WriteStringToFile(mutex *sync.Mutex, fileName string, str string) bool {
	// sync.Mutex ensures oly one goroutine can write to file at a time
	rootDir, err := os.Getwd()
	if err != nil {
		return false
	}

	fileName = path.Join(rootDir, "storage", fileName)
	// opens file for appending, and writing, permissions 0644 allow owner to read and write
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		return false
	}

	defer file.Close()

	mutex.Lock()
	_, err = file.WriteString(str)
	mutex.Unlock()

	return err == nil
}

func WriteBytesToFile(mutex *sync.Mutex, fileName string, data []byte) bool {
	rootDir, err := os.Getwd()
	if err != nil {
		return false
	}

	fileName = path.Join(rootDir, "storage", fileName)
	//  changes size of the file/ clears contents
	if err := os.Truncate(fileName, 0); err != nil {
		return false
	}

	// open file with write permission
	file, err := os.OpenFile(fileName, os.O_WRONLY, 0o644)
	if err != nil {
		return false
	}

	defer file.Close()
	mutex.Lock()
	_, err = file.Write(data)
	mutex.Unlock()

	return err == nil
}
