package daemon

import (
	"context"
	"strings"
	"strconv"
	"github.com/Prominence673/rxemu/internal/ipc"
)

func (d *Daemon) search(req ipc.Request) ipc.Response{
	if len(req.Args) == 0{
		return ipc.Response{
		 OK:    false,
		 Error: "search query is required",
		}
	}
	var query string
	if len(req.Args) == 1{
		query = req.Args[0]
	}else{
		query = strings.Join(req.Args[:len(req.Args)-1], " ")
		query = strings.TrimSpace(query)
	}
	if query == ""{
		return ipc.Response{
		 OK:    false,
		 Error: "search query is required",
		}
	}
	
    parsedLimit, err := strconv.Atoi(req.Args[len(req.Args)-1])
    if err != nil {
        return ipc.Response{
            OK:    false,
            Error: "search limit must be a number",
        }
    }

    limit := parsedLimit

    tracks, err := d.searcher.Search(
        context.Background(),
        query,
        limit,
    )
    
    if err != nil {
        return ipc.Response{
            OK:    false,
            Error: err.Error(),
        }
    }
    d.LastTrack = tracks
    return ipc.Response{
        OK:     true,
        Tracks: tracks,
    }
}