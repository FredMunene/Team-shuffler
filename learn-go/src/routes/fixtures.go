package routes

import (
	"encoding/json"
	"net/http"
	"os"
	"path"

	"learn-go/src/core"
	"learn-go/src/types"
)

func Fixtures(w http.ResponseWriter, r *http.Request) {
	// get working directory
	rootDir, err := os.Getwd()
	if err != nil {
		http.Error(w, "Error creating fixtures", http.StatusInternalServerError)
	}

	// obtain path to file with teams
	filename := path.Join(rootDir, "storage", "teams.txt")

	// access properties of the file
	stats, err := os.Stat(filename)
	if err != nil {
		http.Error(w, "Error creating fixtures", http.StatusInternalServerError)
	}

	// check file is not empty
	if stats.Size() == 0 {
		http.Error(w, "Error creating fixtures", http.StatusInternalServerError)
	}

	// access contents of the file
	content, err := os.ReadFile(filename)
	if err != nil {
		http.Error(w, "Error creating fixtures", http.StatusInternalServerError)
	}

	// create a slice Teams of type struct to store the contents
	teams := []types.Teams{}

	// map the contents into teams variable
	err = json.Unmarshal(content, &teams)
	if err != nil {
		http.Error(w, "Error Creating fixtures", http.StatusInternalServerError)
		return
	}

	// create fixtures
	fixtures := core.Fixtures(teams)

	// map the teams onto json retrievable form
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(fixtures)
}
