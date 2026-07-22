package daemon

import "github.com/Prominence673/rxemu/internal/ipc"

func (d *Daemon) status() ipc.Response {
	return ipc.Response{OK: true, Message: string(d.state.Status)}
}
