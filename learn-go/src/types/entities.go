package types

type Player struct {
	Name string `json:"names"`
	Role string `json:"role"`
}

type Teams struct {
	Name string `json:"name"`
	Stricker string `json:"striker"`
	Defender string `json:"defender"`
}

type Table struct {
	TeamName string `json:"name"`
	Played int `json:"played"`
	Wins int  `json:"wins"`
	Draws int `json:"draws"`
	Losses int `json:"loser"`
	GoalsFor int `json:"goalsfor"`
	GoalsAgainst int`json:"goalsagainst"`
	Points int `json:"points"`
}