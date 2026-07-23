package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

var volumeDownCmd = &cobra.Command{
	Use:   "down",
	Aliases: []string{"decrease", "-"},
	Short: "Lower volume",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		res := changeVolume("vol-down")
		if res.OK{
			fmt.Println(res.Message)
		}else{
			fmt.Println(res.Error)
		}
		return nil
	},
}

func init() {
	volumeCmd.AddCommand(volumeDownCmd)
}