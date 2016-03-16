package server

type Server interface {
	Start() error
	Stop()
	StopNotify() <-chan struct{}
}

type server struct {
	cfg *Config
	//stop notify
	done chan struct{}
	//stop server
	stop chan struct{}
}

func (s *server) Start() (err error) {
	go s.run()
	return
}

func (s *server) run() {
	for {
		select {
		case <-s.stop:
			return
		}
	}
}

func (s *server) Stop() {
	select {
	case s.stop <- struct{}{}:
	case <-s.done:
		return
	}
	<-s.done
}

func (s *server) StopNotify() <-chan struct{} {
	return s.done
}

func NewServer(cfg *Config) (s Server, err error) {
	ts := &server{}
	ts.cfg = cfg
	ts.done = make(chan struct{})
	ts.stop = make(chan struct{})
	s = ts
	return
}
