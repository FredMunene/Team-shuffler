package core

import (
	"learn-go/src/types"
)

func Fixtures(teams []types.Teams){
	n := len(teams)
	if n%2 != 0 {
		teams := append(teams, types.Teams{
			Name: "",
			Stricker: "",
			Defender: "",
		})
	}

	
}
