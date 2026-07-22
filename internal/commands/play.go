package commands

import (
	"fmt"
	"github.com/Prominence673/rxemu/internal/config"
	"github.com/Prominence673/rxemu/internal/ipc"
	"github.com/spf13/cobra"
)

func play(url string) ipc.Response {
	req := ipc.Request{Command: "play", Args: []string{url}}
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

var playCmd = &cobra.Command{
	Use:   "play <url>",
	Short: "play current music",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		res := play(args[0])
		fmt.Println(res)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(playCmd)
}
