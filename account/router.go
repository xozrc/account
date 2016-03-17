package account

import (
	"fmt"
)

import (
	"github.com/gorilla/mux"
)

var _ = fmt.Print

const (
	registerPath    = "/register"
	loginPath       = "/login"
	loginUniqueCode = "/loginUniqueCode"
)

func NewRouter(b Backend) (r *mux.Router) {
	r = &mux.Router{}
	r.HandleFunc(loginPath, Login(b))
	return r
}
