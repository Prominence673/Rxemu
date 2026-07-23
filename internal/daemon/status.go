package daemon

import (
	"fmt"

	"github.com/Prominence673/rxemu/internal/commands"
	"github.com/Prominence673/rxemu/internal/ipc"
	"github.com/Prominence673/rxemu/internal/source"
)

func (d *Daemon) status() ipc.Response {
	if d.CurrentTrack == nil || *d.CurrentTrack == (source.Track{}){
		return ipc.Response{OK: false, Error: "Play a song first"}
	}
	min := commands.FormatDuration(d.CurrentTrack.Duration)
	return ipc.Response{OK: true, Message: fmt.Sprintf("Title %s - Artist %s - Duration %s", d.CurrentTrack.Title, d.CurrentTrack.Artist, min)}
}
