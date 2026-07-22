package main

import (
	"fmt"
	"os"

	"github.com/Prominence673/rxemu/internal/config"
	"github.com/Prominence673/rxemu/internal/daemon"
	"github.com/Prominence673/rxemu/internal/ipc"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	p := daemon.NewPlayer(cfg.MVPsocketPath)

	if err := p.Start(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	defer p.Close()

	d := daemon.New(p)
	ser := ipc.NewServer(cfg.SocketPath, d)

	if err := ser.ListenAndServe(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}
