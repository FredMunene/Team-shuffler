package routes

import (
	"encoding/json"
	"net/http"
	"os"
	"path"

	"learn-go/src/types"
)

func Fixtures(w http.ResponseWriter, r *http.Request) {
	rootDir, err := os.Getwd()
	if err != nil {
		http.Error(w, "Error creating fixtures", http.StatusInternalServerError)
	}
	filename := path.Join(rootDir, "storage", "teams.txt")

	stats, err := os.Stat(filename)
	if err != nil {
		http.Error(w, "Error creating fixtures", http.StatusInternalServerError)
	}
	if stats.Size() == 0 {
		http.Error(w, "Error creating fixtures", http.StatusInternalServerError)
	}

	content, err := os.ReadFile(filename)
	if err != nil {
		http.Error(w, "Error creating fixtures", http.StatusInternalServerError)
	}
	teams := []types.Team{}

	err = json.Unmarshal(content,&teams)
	if err != nil {
		http.Error(w,"Error Creating fixtures", http.StatusInternalServerError)
		return
	}

	fixtures := core.Fixtures(teams)

	w.Header().Set("content-type","application/json")
	json.NewEncoder(w).Encode(fixtures)
}
