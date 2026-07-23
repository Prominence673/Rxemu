package ipc

import "github.com/Prominence673/rxemu/internal/source"

type Request struct {
	Command string   `json:"command"`
	Args    []string `json:"args,omitempty"`
}

type Response struct {
	OK      bool   `json:"ok"`
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
	Tracks []source.Track `json:"tracks,omitempty"`
}
