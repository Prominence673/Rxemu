package daemon

import "github.com/Prominence673/rxemu/internal/ipc"

func (d *Daemon) play(req ipc.Request) ipc.Response {
	if len(req.Args) == 0 || req.Args[0] == "" {
		return ipc.Response{
			OK:    false,
			Error: "URL is required",
		}
	}
	if err := d.player.play(req.Args[0]); err != nil {
		return ipc.Response{
			OK:    false,
			Error: err.Error(),
		}
	}
	d.state.Status = statePlaying
	return ipc.Response{
		OK:      true,
		Message: string(d.state.Status),
	}
}
