package daemon

import "github.com/Prominence673/rxemu/internal/ipc"

func (d *Daemon) stop() ipc.Response {
	if err := d.player.stop(); err != nil {
		return ipc.Response{
			OK:    false,
			Error: err.Error(),
		}
	}
	d.state.Status = stateStopped
	return ipc.Response{OK: true, Message: string(d.state.Status)}
}
