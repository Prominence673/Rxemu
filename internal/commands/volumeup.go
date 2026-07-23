package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

var volumeUpCmd = &cobra.Command{
	Use:   "up",
	Aliases: []string{"increase", "+"},
	Short: "Increase playback volume",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		res := changeVolume("vol-up")
		if res.OK{
			fmt.Println(res.Message)
		}else{
			fmt.Println(res.Error)
		}
		return nil
	},
}

func init() {
	volumeCmd.AddCommand(volumeUpCmd)
}
