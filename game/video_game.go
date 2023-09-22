package game

import "fmt"

type videoGame struct {
	title string
}

func NewVideoGame(title string) *videoGame {
	return &videoGame{
		title: title,
	}
}

func (v *videoGame) Play() {
	fmt.Println("I'm playing a video game!")
}

func (v *videoGame) Title() string {
	return v.title
}
