package daemon

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	accountHandler "github.com/xozrc/account/handler"
	accountServer "github.com/xozrc/account/server"
	"os"
)

var (
	router *mux.Router
)

func init() {
	router = accountHandler.Router
}

func Main() {

	cfg := newConfig()
	stoped, err := startLogin(cfg)
	if err != nil {
		log.Errorf("start login server error %s", err.Error())
		os.Exit(1)
	}

	//stop
	<-stoped
	os.Exit(0)
}

func startLogin(cfg *config) (stopNotify <-chan struct{}, err error) {
	//init config
	srvCfg := accountServer.NewConfig()
	//init server
	s, err := accountServer.NewServer(srvCfg)
	if err != nil {
		return
	}

	//start server
	s.Start()
	stopNotify = s.StopNotify()

	//listen http
	ne := negroni.Classic()

	ne.UseHandler(router)
	listenUrlString := getListenUrlString(cfg)

	go func(n *negroni.Negroni, addr string) {
		n.Run(addr)
	}(ne, listenUrlString)

	return
}

func getListenUrlString(cfg *config) string {
	return fmt.Sprintf("%s:%d", cfg.host, cfg.port)
}
