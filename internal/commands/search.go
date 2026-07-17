package commands

import (
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use: "search <url | title>",
	Short: "search music",
	RunE: func(cmd *cobra.Command, args []string) error{
		if err := cmd.Help(); err != nil{
			return err
		}
		return nil
	},
}

func init(){
	rootCmd.AddCommand(searchCmd)
}
