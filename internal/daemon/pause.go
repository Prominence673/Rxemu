package daemon

import "github.com/Prominence673/rxemu/internal/ipc"

func (d *Daemon) pause() ipc.Response {
	if d.state.Status == statePause {
		if err := d.player.resume(); err != nil {
			return ipc.Response{
				OK:    false,
				Error: err.Error(),
			}
		}

		d.state.Status = statePlaying
	} else {
		if err := d.player.pause(); err != nil {
			return ipc.Response{
				OK:    false,
				Error: err.Error(),
			}
		}

		d.state.Status = statePause
	}

	return ipc.Response{
		OK:      true,
		Message: string(d.state.Status),
	}
}
