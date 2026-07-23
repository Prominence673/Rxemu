package commands

import (
	"github.com/spf13/cobra"
	"github.com/Prominence673/rxemu/internal/config"
	"github.com/Prominence673/rxemu/internal/ipc"
	"fmt"
)
func seekAbs(args []string) ipc.Response {
    cfg, err := config.Load()
    if err != nil {
        return ipc.Response{
            OK:    false,
            Error: err.Error(),
        }
    }

    client := ipc.NewClient(cfg.SocketPath)

    res, err := client.Send(ipc.Request{
        Command: "seek-abs",
        Args: args,
    })
    if err != nil {
        return ipc.Response{
            OK:    false,
            Error: err.Error(),
        }
    }

    return res
}

var seekabsCmd = &cobra.Command{
	Use:   "seekabs <amount>",
	Short: "Advance playback",
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		res := seekAbs(args)
		if res.OK{
			fmt.Println(res.Message)
		}else{
			fmt.Println(res.Error)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(seekabsCmd)
}