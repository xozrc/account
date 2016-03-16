package daemon

type config struct {
	host string
	port int
}

func newConfig() (cfg *config) {
	host := "127.0.0.1"
	port := 3000
	cfg = &config{}
	cfg.host = host
	cfg.port = port
	return
}
