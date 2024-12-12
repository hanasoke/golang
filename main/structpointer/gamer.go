package main

import "fmt"

type Gamer struct {
	Name  string
	Games []string
}

// AddGame
func (gamer *Gamer) AddGame(game string) {
	gamer.Games = append(gamer.Games, game)
}

func main() {
	gamer := Gamer{Name: "Zelda"}

	gamer.AddGame("Borderlands")
	gamer.AddGame("Bioshock")
	gamer.AddGame("Just Dance")
	gamer.AddGame("Residence Evil")
	gamer.AddGame("Jumanji")

	for _, game := range gamer.Games {
		fmt.Println(game)
	}
}
