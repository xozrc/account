package account

import (
	"fmt"
)

import (
	"github.com/gorilla/mux"
	"golang.org/x/net/context"
)

var _ = fmt.Print

const (
	registerPath = "/register.json"
	loginPath    = "/login.json"
)

func NewRouter(ctx context.Context) (r *mux.Router) {
	r = &mux.Router{}
	r.Handle(loginPath, LoginHTTPHandler(ctx))
	return r
}
