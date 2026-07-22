package ipc

import (
	"encoding/json"
	"net"
)

type Client struct {
	socketPath string
}

func NewClient(socketPath string) *Client {
	return &Client{socketPath: socketPath}
}

func (c *Client) Send(req Request) (Response, error) {
	conn, err := net.Dial("unix", c.socketPath)
	if err != nil {
		return Response{}, err
	}
	defer conn.Close()
	if err := json.NewEncoder(conn).Encode(req); err != nil {
		return Response{}, err
	}
	var res Response
	if err := json.NewDecoder(conn).Decode(&res); err != nil {
		return Response{}, err
	}
	return res, nil
}
