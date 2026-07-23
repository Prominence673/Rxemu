package daemon

import (
	"strconv"

	"github.com/Prominence673/rxemu/internal/ipc"
)

func (d *Daemon) volSet(req ipc.Request) ipc.Response{
	amount, err := strconv.ParseFloat(req.Args[0], 64)
	if err != nil{
		return ipc.Response{OK: false, Error: "Invalid value - only decimal/float"}
	}
	if amount < 0 || amount > 100{
		return ipc.Response{OK: false, Error: "Invalid value - (0 - 100)"}
	}
	if err := d.player.volSet(amount); err != nil{
		return ipc.Response{OK: false, Error: err.Error()}	
	}
	return ipc.Response{OK: true, Message: "Volume up"}
}

func (d *Daemon) volUp() ipc.Response{
	if err := d.player.volSet(5); err != nil{
		return ipc.Response{OK: false, Error: err.Error()}	
	}
	return ipc.Response{OK: true, Message: "Volume up"}
}

func (d *Daemon) volDown() ipc.Response{
	if err := d.player.volSet(-5); err != nil{
		return ipc.Response{OK: false, Error: err.Error()}	
	}
	return ipc.Response{OK: true, Message: "Volume down"}
}