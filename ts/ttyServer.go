package gs

import (
	"log"

	"github.com/olahol/melody"
)

type TtyServer struct {
	players map[*melody.Session]*Player
}

func NewServer() *TtyServer {
	return &TtyServer{
		players: make(map[*melody.Session]*Player),
	}
}

func (g *TtyServer) Connect(s *melody.Session) {
	log.Printf("[FINE] %s connected TtyServer.", s.Request.RemoteAddr)
	p := newPlayer()
	g.players[s] = p
}

func (g *TtyServer) Disconnect(s *melody.Session) {
	log.Printf("[FINE] %s disconnected TtyServer.", s.Request.RemoteAddr)
	delete(g.players, s)
}

func (g *TtyServer) Message(s *melody.Session, msg []byte) {
	log.Printf("[DEBUG] gsmsg = %s", msg)

}
