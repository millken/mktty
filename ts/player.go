package gs

import "github.com/olahol/melody"

type Player struct {
	token string
	uid   int
}

func newPlayer() *Player {
	return &Player{}
}

func (p *Player) Connect(s *melody.Session) {
}
