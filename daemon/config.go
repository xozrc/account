package daemon

type config struct {
	Debug   bool   `json:"debug"`
	Host    string `json:"host"`
	Port    int    `json:"port"`
	LogFile string `json:"logFile"`
}

func newConfig() (cfg *config) {
	host := "127.0.0.1"
	port := 3000
	cfg = &config{}
	cfg.Host = host
	cfg.Port = port
	cfg.Debug = true
	cfg.LogFile = "../config/logger.conf"
	return
}
