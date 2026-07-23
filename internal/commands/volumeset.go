package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/Prominence673/rxemu/internal/ipc"
	"github.com/Prominence673/rxemu/internal/config"
)

func setVolume(amount []string) ipc.Response {
    cfg, err := config.Load()
    if err != nil {
        return ipc.Response{
            OK:    false,
            Error: err.Error(),
        }
    }

    client := ipc.NewClient(cfg.SocketPath)

    res, err := client.Send(ipc.Request{
        Command: "vol-set",
        Args: amount,
    })
    if err != nil {
        return ipc.Response{
            OK:    false,
            Error: err.Error(),
        }
    }

    return res
}

var volumeSetCmd = &cobra.Command{
	Use:   "set <amount>",
	Short: "set volume beetween 0 - 100",
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		res := setVolume(args)
		if res.OK{
			fmt.Println(res.Message)
		}else{
			fmt.Println(res.Error)
		}
		return nil
	},
}

func init() {
	volumeCmd.AddCommand(volumeSetCmd)
}