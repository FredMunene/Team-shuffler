package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"learn-go/src/core"
	"learn-go/src/types"
)

func Register(mutex *sync.Mutex) http.HandlerFunc {
	// sync.Mutex ensures oly one goroutine can write to file at a time
	return func(w http.ResponseWriter, r *http.Request) {
		// extracting name and role from request
		role := r.FormValue("role")
		name := r.FormValue("name")
		fmt.Println(role + " " + name) // prints value to console
		if role == "" || name == "" {
			// 400 status code if parameter is missing
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		// variable to store name and role
		line := fmt.Sprintf("\n%s : %s", role, name)

		successWrite := false

		// use a goroutine to avoid blocking other requests while satisfying an IO bound process
		// it writes a string to a file in a concurrent manner using a mutex for thread-safe access
		go func() {
			successWrite = core.WriteStringToFile(mutex, "players.txt", line)
		}()

		if successWrite {
			http.Error(w, "Error registering user", http.StatusInternalServerError)
			return
		}

		player := types.Player{Role: role, Name: name}
		w.Header().Set("content-type", "application/json")
		json.NewEncoder(w).Encode(player)
	}
}
