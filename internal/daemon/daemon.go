package daemon

import "github.com/Prominence673/rxemu/internal/ipc"

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
}

func New(p *Player) *Daemon {
	return &Daemon{state: State{Status: stateStopped}, player: p}
}

func (d *Daemon) Handle(req ipc.Request) ipc.Response {
	switch req.Command {
	case "status":
		return d.status()
	case "stop":
		return d.stop()
	case "play":
		return d.play(req)
	case "pause":
		return d.pause()
	default:
		return ipc.Response{OK: false, Message: "Invalid command"}
	}
}
