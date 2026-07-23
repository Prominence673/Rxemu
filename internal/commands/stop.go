package commands

import (
	"fmt"
	"github.com/Prominence673/rxemu/internal/config"
	"github.com/Prominence673/rxemu/internal/ipc"
	"github.com/spf13/cobra"
)

func stop() ipc.Response {
	req := ipc.Request{Command: "stop"}
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

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "stop current music",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		res := stop()
		if res.OK{
			fmt.Println(res.Message)
		}else{
			fmt.Println(res.Error)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
