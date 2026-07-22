package ipc

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"os"
)

type Handler interface {
	Handle(req Request) Response
}

type Server struct {
	socketPath string
	handler    Handler
}

func NewServer(socketPath string, handler Handler) *Server {
	return &Server{socketPath: socketPath, handler: handler}
}

func (s *Server) ListenAndServe() error {
	if err := s.prepareSocket(); err != nil {
		return err
	}

	listen, err := net.Listen("unix", s.socketPath)
	if err != nil {
		return fmt.Errorf("listen on unix socket : %w", err)
	}

	defer listen.Close()

	defer os.Remove(s.socketPath)

	for {
		conn, err := listen.Accept()
		if err != nil {
			return fmt.Errorf("accept connection: %w", err)
		}
		if err := s.handleConn(conn); err != nil {
			fmt.Fprintln(os.Stderr, "IPC connection error:", err)
		}
	}
}

func (s *Server) prepareSocket() error {
	conn, err := net.Dial("unix", s.socketPath)
	if err == nil {
		conn.Close()
		return errors.New("another socket daemon is already running")
	}
	if err := os.Remove(s.socketPath); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("remove stale socket: %w", err)
	}
	return nil
}

func (s *Server) handleConn(conn net.Conn) error {
	defer conn.Close()

	var req Request
	if err := json.NewDecoder(conn).Decode(&req); err != nil {
		return fmt.Errorf("decode request: %w", err)
	}
	res := s.handler.Handle(req)

	if err := json.NewEncoder(conn).Encode(res); err != nil {
		return fmt.Errorf("encode response: %w", err)
	}

	return nil
}
