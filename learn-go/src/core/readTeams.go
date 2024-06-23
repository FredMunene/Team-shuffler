package core

import (
	"bufio"
	"log"
	"os"
	"path"
	"strings"

	
)

func ReadTeams(filename string) ([]string, []string, error) {
	// get current working directory
	rootDir, err := os.Getwd()
	if err != nil {
		return nil, nil, err
	}

	filename = path.Join(rootDir, "storage", filename)
	// open file 
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()
	// variables to store names of strickers and defenders
	var strickers []string
	var defenders []string

	// scan file
	scanner := bufio.NewScanner(file)
	// scan file line by line
	for scanner.Scan() {
		line := scanner.Text()
		// spliting line based on space
		words := strings.Split(line, " ")
		// if first word is stricker, append other words from second one to slice Strickers
		if strings.ToLower(words[0]) == "stricker:" {
			strickers = append(strickers, strings.Join(words[1:]," "))
		} else if strings.ToLower(words[0]) == "defender:" {
			defenders  = append(defenders, strings.Join(words[1:], " "))
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
	// if len(strickers) == 0 ||
	// if len(defenders) == 0 ||
	// if len(defenders) != len(strickers)
	// return strickers, defenders, fmt.Errorf("number of strickers and defenders are not equal")

	return strickers, defenders, nil
}
