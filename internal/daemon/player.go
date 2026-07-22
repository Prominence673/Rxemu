package daemon

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"os/exec"
	"time"
)

type Player struct {
	command    *exec.Cmd
	socketPath string
}

type MPVRequest struct {
	Command []any `json:"command"`
}

type MPVResponse struct {
	Event string `json:"event"`
	Error string `json:"error"`
	Data  any    `json:"data,omitempty"`
}

func NewPlayer(socketPath string) *Player {
	cmd := exec.Command(
		"mpv",
		"--idle=yes",
		"--no-video",
		"--no-config",
		"--input-ipc-server="+socketPath,
	)
	return &Player{command: cmd, socketPath: socketPath}
}

func (p *Player) Start() error {
	if err := os.Remove(p.socketPath); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("remove old mpv socket: %w", err)
	}
	if err := p.command.Start(); err != nil {
		return fmt.Errorf("start mpv: %w", err)
	}
	return p.waitForSocket(2 * time.Second)
}

func (p *Player) send(command []any) error {
	conn, err := net.DialTimeout(
		"unix",
		p.socketPath,
		2*time.Second,
	)
	if err != nil {
		return fmt.Errorf("connect to mpv: %w", err)
	}
	defer conn.Close()

	if err := conn.SetDeadline(time.Now().Add(5 * time.Second)); err != nil {
		return fmt.Errorf("set mpv deadline: %w", err)
	}

	req := MPVRequest{
		Command: command,
	}

	if err := json.NewEncoder(conn).Encode(req); err != nil {
		return fmt.Errorf("send command to mpv: %w", err)
	}

	decoder := json.NewDecoder(conn)

	for {
		var res MPVResponse

		if err := decoder.Decode(&res); err != nil {
			return fmt.Errorf("read mpv response: %w", err)
		}

		if res.Event != "" {
			continue
		}

		if res.Error != "success" {
			return fmt.Errorf("mpv command failed: %s", res.Error)
		}

		return nil
	}
}

func (p *Player) play(url string) error {
	if err := p.send([]any{
		"loadfile",
		url,
		"replace",
	}); err != nil {
		return fmt.Errorf("play file: %w", err)
	}

	return nil
}

func (p *Player) pause() error {
	if err := p.send([]any{
		"set_property",
		"pause",
		true,
	}); err != nil {
		return fmt.Errorf("pause playback: %w", err)
	}

	return nil
}

func (p *Player) resume() error {
	if err := p.send([]any{
		"set_property",
		"pause",
		false,
	}); err != nil {
		return fmt.Errorf("resume playback: %w", err)
	}

	return nil
}

func (p *Player) stop() error {
	if err := p.send([]any{"stop"}); err != nil {
		return fmt.Errorf("stop playback: %w", err)
	}

	return nil
}

func (p *Player) waitForSocket(timeout time.Duration) error {
	deadline := time.Now().Add(timeout)

	for time.Now().Before(deadline) {
		if _, err := os.Stat(p.socketPath); err == nil {
			return nil
		}
		time.Sleep(50 * time.Millisecond)
	}

	_ = p.command.Process.Kill()
	_ = p.command.Wait()

	return fmt.Errorf("mpv did not create socket within %s", timeout)
}

func (p *Player) Close() error {
	if p.command.Process == nil {
		return nil
	}
	if err := p.command.Process.Kill(); err != nil {
		return fmt.Errorf("stop mpv: %w", err)
	}
	if err := p.command.Wait(); err != nil {
		return nil
	}
	return nil
}
