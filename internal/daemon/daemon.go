package daemon

import (
	"github.com/Prominence673/rxemu/internal/ipc"
	"github.com/Prominence673/rxemu/internal/source"
)

type PlayStatus string

const (
	stateStopped PlayStatus = "stopped"
	statePlaying PlayStatus = "playing"
	statePause   PlayStatus = "paused"
)

type State struct {
	Status PlayStatus
}

type Daemon struct {
	state  State
	player *Player
	searcher source.Searcher
	CurrentTrack *source.Track
	LastTrack []source.Track
}

func New(p *Player, s source.Searcher) *Daemon {
	return &Daemon{
		state: State{Status: stateStopped}, 
		player: p,
		searcher: s,
	}
}
func (d *Daemon) WatchPlayer() {
	for event := range d.player.Events() {
		switch event.Event {
		case "file-loaded":
			d.state.Status = statePlaying

		case "end-file":
			d.state.Status = stateStopped
		}
	}
}
func (d *Daemon) Handle(req ipc.Request) ipc.Response {
	switch req.Command {
	case "status":
		return d.status()
	case "stop":
		return d.stop()
	case "play":
		return d.play(req)
	case "playurl":
		return d.playURL(req)
	case "pause":
		return d.pause()
	case "search":
		return d.search(req)
	default:
		return ipc.Response{OK: false, Message: "Invalid command"}
	}
}
