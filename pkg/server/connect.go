package server

import (
	"context"

	"github.com/google/uuid"
	pb "github.com/movaua/rock-paper-scissors/pkg/rps"
)

// Connect connects a player to the game.
// Request Player must have Name set (Id is ignored).
// Response Player is assigned Id.
func (g *game) Connect(ctx context.Context, p *pb.Player) (*pb.Player, error) {
	r := connectRequest{player: p, res: make(chan connectResponse)}
	g.connectRequests <- r
	res := <-r.res
	return res.player, res.err
}

type connectRequest struct {
	player *pb.Player
	res    chan connectResponse
}

type connectResponse struct {
	player *pb.Player
	err    error
}

func (g *game) handleConnect(r connectRequest) {
	defer close(r.res)

	if r.player.GetName() == "" {
		r.res <- connectResponse{err: errEmptyName}
		return
	}
	if player := g.findPlayerByName(r.player.GetName()); player != nil {
		r.res <- connectResponse{err: errConnected(r.player.GetName())}
		return
	}
	if g.isStarted() {
		r.res <- connectResponse{err: errStarted}
		return
	}

	player := &pb.Player{
		Id:   uuid.New().String(),
		Name: r.player.GetName(),
	}
	g.players[player.Id] = player
	r.res <- connectResponse{player: player}

	for _, notifyPlayerConnected := range g.notifyPlayerConnectedChans {
		notifyPlayerConnected <- player
	}

	if len(g.players) > 1 {
		g.start()
	}
}

func (g *game) findPlayerByName(name string) *pb.Player {
	for _, player := range g.players {
		if player.GetName() == name {
			return player
		}
	}
	return nil
}
