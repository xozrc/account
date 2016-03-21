package account

import (
	"github.com/xozrc/account/types"
	"golang.org/x/net/context"
)

const (
	accountBackendKey = "AccountBackend"
)

type AccountBackend interface {
	Login(at types.AccountType, secondId string, uniqueCode string) (acc *types.Account, err error)
}

type accountBackend struct {
}

func (ab *accountBackend) Login(at types.AccountType, secondId string, uniqueCode string) (acc *types.Account, err error) {
	return
}

func NewAccountBackend() AccountBackend {
	b := &accountBackend{}
	return b
}

func AccountBackendInContext(ctx context.Context) (ab AccountBackend) {
	tab := ctx.Value(accountBackendKey)
	ab, _ = tab.(AccountBackend)
	return

}

func WithAccountBackend(ctx context.Context, ab AccountBackend) context.Context {
	tc := context.WithValue(ctx, accountBackendKey, ab)
	return tc
}
