package game

type store struct {
	games []Game
}

func NewStore() *store {
	return &store{}
}

func (s *store) Games() []Game {
	return s.games
}

func (s *store) AddGame(game Game) {
	s.games = append(s.games, game)
}
