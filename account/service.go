package account

import (
	"github.com/xozrc/account/types"
	"golang.org/x/net/context"
)

const (
	AccountServiceKey = "AccountService"
)

func NewAccountService(b Backend) (as *AccountService) {
	as = &AccountService{}
	as.backend = b
	return as
}

type AccountService struct {
	backend Backend
}

func (as *AccountService) Login(at types.AccountType, secondId string, uniqueCode string) (acc *types.Account, err error) {

	switch at {
	case types.Visitor:
		acc = as.backend.AccountByUniqueCode(uniqueCode)
		break
	default:
		acc = as.backend.AccountByTypeAndSecondId(int(at), secondId)
	}

	//no exist user
	if acc == nil {
		//new user
	}

	return
}

func WithAccountService(ctx context.Context, as *AccountService) context.Context {
	return context.WithValue(ctx, AccountServiceKey, as)
}

func AccountServiceInContext(ctx context.Context) (as *AccountService) {
	tas := ctx.Value(AccountServiceKey)
	as, _ = tas.(*AccountService)
	return
}
