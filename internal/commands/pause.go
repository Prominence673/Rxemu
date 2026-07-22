package commands

import (
	"fmt"

	"github.com/Prominence673/rxemu/internal/config"
	"github.com/Prominence673/rxemu/internal/ipc"
	"github.com/spf13/cobra"
)

func pause() ipc.Response {
	req := ipc.Request{Command: "pause"}
	cfg, err := config.Load()
	if err != nil {
		return ipc.Response{OK: false, Error: err.Error()}
	}
	client := ipc.NewClient(cfg.SocketPath)
	res, err := client.Send(req)
	if err != nil {
		return ipc.Response{OK: false, Error: err.Error()}
	}
	return res
}

var pauseCmd = &cobra.Command{
	Use:   "pause",
	Short: "pause current music",
	RunE: func(cmd *cobra.Command, args []string) error {
		res := pause()
		fmt.Println(res)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(pauseCmd)
}
