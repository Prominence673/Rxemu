package commands

import (
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use: "status",
	Short: "music status",
	RunE: func(cmd *cobra.Command, args []string) error{
		if err := cmd.Help(); err != nil{
			return err
		}
		return nil
	},
}

func init(){
	rootCmd.AddCommand(statusCmd)
}
