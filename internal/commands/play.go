package commands

import (
	"github.com/spf13/cobra"
)

var playCmd = &cobra.Command{
	Use: "play",
	Short: "play current music",
	RunE: func(cmd *cobra.Command, args []string) error{
		if err := cmd.Help(); err != nil{
			return err
		}
		return nil
	},
}

func init(){
	rootCmd.AddCommand(playCmd)
}
