package daemon

import (
	"github.com/Prominence673/rxemu/internal/ipc"
	"github.com/Prominence673/rxemu/internal/source"
)

func (d *Daemon) stop() ipc.Response {
	if err := d.player.stop(); err != nil {
		return ipc.Response{
			OK:    false,
			Error: err.Error(),
		}
	}
	d.state.Status = stateStopped
	d.CurrentTrack = &source.Track{}
	return ipc.Response{OK: true, Message: "■ Stopped"}
}
