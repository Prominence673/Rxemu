package main

import (
	"fmt"
	"os"
	"github.com/Prominence673/rxemu/internal/config"
	"github.com/Prominence673/rxemu/internal/daemon"
	"github.com/Prominence673/rxemu/internal/ipc"
	"github.com/Prominence673/rxemu/internal/source"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	p := daemon.NewPlayer(cfg.MVPsocketPath)
	y := source.NewYouTube()

	if err := p.Start(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	go func() {
    if err := p.ListenEvents(); err != nil {
        fmt.Fprintln(os.Stderr, "mpv event listener:", err)
    }
	}()
	defer p.Close()

	d := daemon.New(p, y)
	go d.WatchPlayer()
	ser := ipc.NewServer(cfg.SocketPath, d)

	if err := ser.ListenAndServe(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}
