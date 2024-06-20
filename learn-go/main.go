package main

import (
	"fmt"
	"learn-go/src/core"
	"os"
	"path"
)

func main() {
	url := "https://raw.githubusercontent.com/FredMunene/ascii-art/main/ascii-art-fs/shadow.txt"
	filename := "./storage/ascii.art"

	core.FetchContent(url, filename)

	fmt.Println(os.Getwd())
}


func CreateOrOpenFile(filename string)(string, bool){
	rootDir, err := os.Getwd()
	if err != nil {
		return "Error accessing directory...", false
	}

	filename = path.Join(rootDir,"storage",filename)

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Sprintf("Error to create \n",filename),false
	}

	defer file.Close()
	return filename,true
}