package config

type Config struct {
	SocketPath    string
	MVPsocketPath string
	HistoryPath   string
	Volume        int
}

func Default() Config {
	return Config{
		SocketPath:    "/tmp/rxemu.sock",
		MVPsocketPath: "/tmp/mvprxemu.sock",
		HistoryPath:   "/tmp/history-rxemu.sock",
		Volume:        100,
	}
}

func Load() (Config, error) {
	return Default(), nil
}
