package core

import (
	"math/rand"
	"time"
)


func Shuffle(players []string) []string{
	// initialize a random number generator
	// seeded with current time in nanoseconds
	// random numbers are different each time the program runs
	rand.New(rand.NewSource(time.Now().UnixNano()))
	rand.Shuffle(len(players),func(i, j int) {
		players[i],players[j] = players[j],players[i]
	})
	return players

}