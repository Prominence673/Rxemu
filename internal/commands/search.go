package commands

import (
	"fmt"
	"github.com/Prominence673/rxemu/internal/config"
	"github.com/Prominence673/rxemu/internal/ipc"
	"github.com/spf13/cobra"
	"strconv"
	"time"
)

var limit int

func FormatDuration(seconds float64) string {
	duration := time.Duration(seconds) * time.Second

	minutes := int(duration.Minutes())
	remainingSeconds := int(duration.Seconds()) % 60

	return fmt.Sprintf(
		"%02d:%02d",
		minutes,
		remainingSeconds,
	)
}

func search(args []string) ipc.Response{
	args = append(args, strconv.Itoa(limit))
	req := ipc.Request{ Command: "search", Args: args}
	cfg, err := config.Load()
	if err != nil{
		return ipc.Response{OK: false, Error: err.Error()}
	}
	client := ipc.NewClient(cfg.SocketPath)
	res, err := client.Send(req)
	if err != nil{
		return ipc.Response{OK: false, Error: err.Error()}
	}
	return res
}

type ViewTrack struct{
	ID int
	Title string
	Artist string
	Duration string
}

var searchCmd = &cobra.Command{
	Use:   "search <url | title>",
	Short: "search music",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		res := search(args)
		result := res.Tracks
		var view []ViewTrack
		for i, t := range result{
			view = append(view, ViewTrack{
				ID: i + 1,
				Title: t.Title,
				Artist: t.Artist,
				Duration: FormatDuration(t.Duration),
			})
		}
		for _, t := range view{
			fmt.Printf("%d. Title: %s Artist: %s Duration: %s\n", t.ID, t.Title, t.Artist, t.Duration)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().IntVarP(&limit, "limit", "l", 5, "results limit by default 5")
}
