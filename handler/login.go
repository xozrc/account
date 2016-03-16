package handler

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var _ = fmt.Print

const (
	loginPath = "/login"
)

const (
	Visitor int = iota
	Facebook
)

var (
	Router *mux.Router
)

func init() {
	Router = &mux.Router{}
	Router.HandleFunc(loginPath, Login)
}

func Login(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(rw, "hello")
}
