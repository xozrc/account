package account

import (
	"fmt"
	"net/http"
)

import (
	"github.com/xozrc/account/types"
	"github.com/xozrc/pkg/httputils"
	"github.com/xozrc/rest"
	"golang.org/x/net/context"
)

var _ = fmt.Print

//request form
type LoginForm struct {
	AccountType int    `json, form:"accountType"`
	SecondId    string `json,form:"secondId"`
	UniqueCode  string `json,form:"uniqueCode"`
}

//request result
type LoginReturnObj struct {
	AccountType int    `json:"accountType"`
	SecondId    string `json:"secondId"`
	UniqueCode  string `json:"uniqueCode"`
}

func loginReturnObjForAccount(acc *types.Account) (lrj *LoginReturnObj) {
	lrj = &LoginReturnObj{}
	lrj.AccountType = acc.AccountType
	lrj.SecondId = acc.SecondId
	lrj.UniqueCode = acc.UniqueCode
	return
}

//rest login handlers
func Login(ctx context.Context, rw http.ResponseWriter, req *http.Request, next rest.ContextHandler) {

	lfi := rest.FormInContext(ctx)
	if lfi == nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	lf, _ := lfi.(*LoginForm)
	if lf == nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	ab := AccountBackendInContext(ctx)
	if ab == nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	acc, err := ab.Login(types.AccountType(lf.AccountType), lf.SecondId, lf.UniqueCode)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	//process acc
	returnObj := loginReturnObjForAccount(acc)
	httputils.WriteJSON(rw, http.StatusOK, returnObj)
}

func LoginHTTPHandler(ctx context.Context) http.Handler {
	return rest.HTTPHandlers(ctx, rest.BindFormHandler(&LoginForm{}), rest.MiddlewareContextHandlerFunc(Login))
}
