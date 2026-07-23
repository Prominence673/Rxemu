package commands

import (
	"github.com/spf13/cobra"
	"github.com/Prominence673/rxemu/internal/config"
	"github.com/Prominence673/rxemu/internal/ipc"
)

func changeVolume(command string) ipc.Response {
    cfg, err := config.Load()
    if err != nil {
        return ipc.Response{
            OK:    false,
            Error: err.Error(),
        }
    }

    client := ipc.NewClient(cfg.SocketPath)

    res, err := client.Send(ipc.Request{
        Command: command,
    })
    if err != nil {
        return ipc.Response{
            OK:    false,
            Error: err.Error(),
        }
    }

    return res
}

var volumeCmd = &cobra.Command{
    Use:     "volume",
    Aliases: []string{"vol"},
    Short:   "Control playback volume",
}

func init() {
    rootCmd.AddCommand(volumeCmd)
}