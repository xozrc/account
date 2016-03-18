package account

import (
	"fmt"
	"net/http"
)

import (
	"github.com/xozrc/account/types"
	"golang.org/x/net/context"
)

var _ = fmt.Print

type LoginForm struct {
	AccountType int
	SecondId    string
	UniqueCode  string
}

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

//rest login handler
func Login(ctx context.Context, rw http.ResponseWriter, req *http.Request) (code int, result interface{}, err error) {

	lf := &LoginForm{}
	lf.AccountType = 1
	lf.SecondId = "asd"
	lf.UniqueCode = "as"

	as := AccountServiceInContext(ctx)
	if as == nil {
		return http.StatusInternalServerError, nil, nil
	}

	acc, err := as.Login(types.AccountType(lf.AccountType), lf.SecondId, lf.UniqueCode)
	if err != nil {
		return http.StatusInternalServerError, nil, nil
	}

	//process acc
	returnObj := loginReturnObjForAccount(acc)

	return http.StatusOK, returnObj, nil
}
