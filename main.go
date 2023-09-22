package main

import (
	"fmt"
	"main/game"
)

func main() {
	vg := game.NewVideoGame("Pac-Man")
	bg := game.NewBoardGame("Catan")
	store := game.NewStore()
	store.AddGame(vg)
	store.AddGame(bg)
	for _, g := range store.Games() {
		fmt.Println(g.Title())
		g.Play()
	}
}
