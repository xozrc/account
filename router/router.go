package router

import (
	"fmt"
)

import (
	"github.com/gorilla/mux"
	"github.com/xozrc/account/types"
)

var _ = fmt.Print

const (
	loginPath = "/login"
)

type Backend interface {
	Register(name string, pass string) (*types.Account, error)
	Login(id string) (*types.Account, error)
	LoginUniqueCode(uniqueCode string) (*types.Account, error)
}

func NewRouter(b Backend) (r *mux.Router) {
	r = &mux.Router{}
	r.HandleFunc(loginPath, Login(b))
	return r
}
