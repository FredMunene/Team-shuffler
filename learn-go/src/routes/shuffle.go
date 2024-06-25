package routes

import (
	"encoding/json"
	"learn-go/src/core"
	"learn-go/src/types"
	"log"
	"net/http"
	"sync"
)



func Shuffle(mutex *sync.Mutex) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		strikers, defenders, err := core.ReadTeams("players.txt")
		if err != nil {
			log.Fatal(err)
		}
		defenders = core.Shuffle(defenders)
		teams:= []types.Teams{}
		for i := 0; i < len(strikers); i++ {
			teamName := core.GenerateString(strikers[i], defenders[i])
			teams = append(teams, types.Teams{Name: teamName, Stricker: strikers[i], Defender: defenders[i]})
		}
		// content is a JSON encoding of teams 
		content, err := json.Marshal(teams)
		if err != nil {
			http.Error(w, "Erroe creating teams json structure", http.StatusInternalServerError)
			return
		}
		// write the content into our file
		if ok := core.WriteBytesToFile(mutex, "teams.csv", content); ok {
			http.Error(w, "Eror creating fixtures", http.StatusInternalServerError)
			return
		}
		w.Header().Set("content-type", "application-json")
		json.NewEncoder(w).Encode(teams)
	}

}