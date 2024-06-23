package routes

import (
	"encoding/json"
	"net/http"

	"learn-go/src/core"
	"learn-go/src/types"
)




func Players(w http.ResponseWriter, r *http.Request){

	strickers, defenders, err := core.ReadTeams("players.txt")
	if err != nil {
		http.Error(w, "Error reading users", http.StatusInternalServerError)
		return
	}

	allPlayers := []types.Player{}
	for _, player := range defenders{
		allPlayers = append(allPlayers, types.Player{Role: "Defender", Name: player})
		
	}
	for _, player := range strickers{
		allPlayers = append(allPlayers, types.Player{Role: "Stricker", Name: player})
		
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(allPlayers)

}

func PlayerlistHandler(w http.ResponseWriter, r *http.Request){

	http.ServeFile(w,r,"templates/players.html")

}