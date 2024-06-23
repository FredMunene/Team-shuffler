package db

import (
	"fmt"
	"os"
	"strconv"

	"learn-go/src/types"

	"github.com/go-gota/gota/dataframe"
)

func GetPlayers() []types.Table {
	file, err := os.Open("storage/team.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	// ReadCSV reads a CSV file from a io.Reader
	csvDf := dataframe.ReadCSV(file)
	// sort data by points
	sorted := csvDf.Arrange(
		dataframe.RevSort("points"),
	)

	// create a struct of type Table
	table := make([]types.Table, 0)
	// convert the dataframe to a slice of strings
	records := sorted.Records()
	// loop through the slice 
	for i, row := range records {
		// skip the header wow
		if i == 0 {
			continue
		}
		table = append(table, types.Table{
			TeamName:     row[0],
			Played:       toInt(row[1]),
			Wins:         toInt(row[2]),
			Draws:        toInt(row[3]),
			Losses:       toInt(row[4]),
			GoalsFor:     toInt(row[5]),
			GoalsAgainst: toInt(row[6]),
			Points:       toInt(row[7]),
		})
	}
	return table
}

func toInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return num
}
