package daemon

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	accountRouter "github.com/xozrc/account/router"
	accountServer "github.com/xozrc/account/server"
	"os"
	"runtime"
)

var (
	router *mux.Router
)

func init() {
	//init router
	router = accountRouter.NewRouter(nil)

}

func Main() {

	cfg := newConfig()

	//set cpu
	maxProcs := runtime.GOMAXPROCS(0)
	log.Infof("setting maximum number of CPUs to %d, total number of available CPUs is %d", maxProcs, runtime.NumCPU())

	stoped, err := startLogin(cfg)
	if err != nil {
		log.Errorf("start login server error %s", err.Error())
		os.Exit(1)
	}

	//todo:register interrupt

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
