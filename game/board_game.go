package game

import "fmt"

type boardGame struct {
	title string
}

func NewBoardGame(title string) *boardGame {
	return &boardGame{
		title: title,
	}
}

func (b *boardGame) Play() {
	fmt.Println("I'm playing a board game!")
}

func (b *boardGame) Title() string {
	return b.title
}
