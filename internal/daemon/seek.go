package daemon

import (
	"fmt"
	"strconv"
	"github.com/Prominence673/rxemu/internal/ipc"
)

func (d *Daemon) seek(req ipc.Request) ipc.Response{
	amount, err := strconv.Atoi(req.Args[0])
	if err != nil{
		return ipc.Response{OK: false, Error: "Invalid value - only decimal"}
	}
	if err := d.player.seek(amount); err != nil{
		return ipc.Response{OK: false, Error: err.Error()}	
	}
	return ipc.Response{OK: true, Message: fmt.Sprintf("avanced %d", amount)}
}

func (d *Daemon) seekAbs(req ipc.Request) ipc.Response{
	amount, err := strconv.Atoi(req.Args[0])
	if err != nil{
		return ipc.Response{OK: false, Error: "Invalid value - only decimal"}
	}
	if err := d.player.seekAbs(amount); err != nil{
		return ipc.Response{OK: false, Error: err.Error()}	
	}
	return ipc.Response{OK: true, Message: fmt.Sprintf("avanced %d", amount)}
}
