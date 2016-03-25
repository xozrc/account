package daemon

import (
	"fmt"
	"os"
	"runtime"
)

import (
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/negroni"

	"github.com/gorilla/mux"
	"github.com/xozrc/account/account"
	accountServer "github.com/xozrc/account/server"
)

var (
	router *mux.Router
)

func init() {
	//init router
	router = account.NewRouter(nil)

}

func Main() {
	//todo: parse config
	//1.from cfg file
	//2.from cmd line
	cfg := newConfig()

	err := setupLogging("")

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
	srvCfg := accountServer.NewServerConfig()
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
	return fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
}

func setupLogging(logFile string) error {
	// if mateConf, err := logrus_mate.LoadLogrusMateConfig(logFile); err != nil {
	// 	return
	// } else {
	// 	if newMate, err := logrus_mate.NewLogrusMate(mateConf); err != nil {
	// 		return
	// 	} else {
	// 		newMate.Logger("mike").Errorln("I am mike in new logrus mate")
	// 	}
	// }
	return nil
}
