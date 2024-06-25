package core

import (
	"learn-go/src/types"
)

func Fixtures(teams []types.Teams)[][]string{
	n := len(teams)
	// incase number of teams is not even, add an empty team
	if n%2 != 0 {
		teams = append(teams, types.Teams{Name: "",	Stricker: "",Defender: ""})
		n++
	}
	// var to hold teams
	fixtures := [][]string{}
	// every team to face every other team??
	for round := 1; round < n-1; round++ {
		// variable to hold all matches??
		match_list := []string{}
		// number of matches will be less number of teams
		for match := 0; match < n/2; match++ {
			// home is 
			home := (round + match) % (n-1)
			// away 
			away := (n - 1 - match + round) % (n-1)
			if match == 0 {
				away = n - 1
			}
			if !(teams[away].Name == "" || teams[home].Name == ""){
				match_list = append(match_list, teams[home].Name + "vs" + teams[away].Name + "\n")
			}
		}
		fixtures = append(fixtures, match_list)
	}
	return fixtures
	
}
