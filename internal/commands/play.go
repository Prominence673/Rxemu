package commands

import (
	"fmt"
	"github.com/Prominence673/rxemu/internal/config"
	"github.com/Prominence673/rxemu/internal/ipc"
	"github.com/spf13/cobra"
)

func play(arg string, url bool) ipc.Response {
	var req ipc.Request
	if url{
		req = ipc.Request{Command: "playurl", Args: []string{arg}}
	}else{
		req = ipc.Request{Command: "play", Args: []string{arg}}
	}
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
	Use:   "play <id>",
	Short: "play current music",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var res ipc.Response
		if v,_ := cmd.Flags().GetBool("url"); v{
			res = play(args[0], true)
		} else{
			res = play(args[0], false)
		}
		if res.OK{
			fmt.Println(res.Message)
		}else{
			fmt.Println(res.Error)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(playCmd)
	playCmd.Flags().BoolP("url", "u", false, "play <url> -u")
}
