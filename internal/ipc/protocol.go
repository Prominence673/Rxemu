package ipc

type Request struct {
	Command string   `json:"command"`
	Args    []string `json:"query,omitempty"`
}

type Response struct {
	OK      bool   `json:"ok"`
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}
