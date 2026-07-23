package daemon

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Prominence673/rxemu/internal/commands"
	"github.com/Prominence673/rxemu/internal/ipc"
)

func (d *Daemon) play(req ipc.Request) ipc.Response{
	if len(req.Args) == 0 || req.Args[0] == "" {
		return ipc.Response{
			OK:    false,
			Error: "ID is required",
		}
	}
	if len(d.LastTrack) == 0{
		return ipc.Response{
			OK:    false,
			Error: "First search a song",
		}
	}
	
	id, err := strconv.Atoi(req.Args[0])
	if err != nil{
		return ipc.Response{
			OK:    false,
			Error: "Invalid ID",
		}
	}
	if id < 1 || id > len(d.LastTrack) {
    return ipc.Response{
        OK:    false,
        Error: "track number is out of range",
    }
	}
	if err := d.player.play(d.LastTrack[id-1].URL); err != nil {
		return ipc.Response{
			OK:    false,
			Error: err.Error(),
		}
	}
	d.state.Status = statePlaying
	d.CurrentTrack = &d.LastTrack[id-1]
	min := commands.FormatDuration(d.CurrentTrack.Duration)
	return ipc.Response{
		OK:      true,
		Message: fmt.Sprintf("Playing ▶︎ %s - %s - %s", d.CurrentTrack.Title, d.CurrentTrack.Artist, min),
	}
} 

func (d *Daemon) playURL(req ipc.Request) ipc.Response {
    if len(req.Args) == 0 || req.Args[0] == "" {
        return ipc.Response{
            OK:    false,
            Error: "URL is required",
        }
    }

    track, err := d.searcher.Resolve(
        context.Background(),
        req.Args[0],
    )
    if err != nil {
        return ipc.Response{
            OK:    false,
            Error: err.Error(),
        }
    }

    if err := d.player.play(track.URL); err != nil {
        return ipc.Response{
            OK:    false,
            Error: err.Error(),
        }
    }

    d.CurrentTrack = &track
    d.state.Status = statePlaying
    min := commands.FormatDuration(d.CurrentTrack.Duration)
    return ipc.Response{
        OK:      true,
        Message: fmt.Sprintf("Playing ▶︎ %s - %s - %s", d.CurrentTrack.Title, d.CurrentTrack.Artist, min),
    }
}
