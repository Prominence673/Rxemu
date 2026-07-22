package daemon

import (
	"github.com/Prominence673/rxemu/internal/ipc"
)

func (d *Daemon) pause() ipc.Response {
	switch d.state.Status{
		case statePause:
			if err := d.player.resume(); err != nil {
				return ipc.Response{
					OK:    false,
					Error: err.Error(),
				}
			}
			d.state.Status = statePlaying
		case statePlaying:
			if err := d.player.pause(); err != nil {
				return ipc.Response{
					OK:    false,
					Error: err.Error(),
				}
			}
			d.state.Status = statePause
		default:
			return ipc.Response{
				OK:    false,
				Error: "Invalid State",
			}
	}
	return ipc.Response{
		OK:      true,
		Message: string(d.state.Status),
	}
}
