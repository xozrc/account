package account

import (
	"github.com/xozrc/account/types"
)

type Backend interface {
	AccountByTypeAndSecondId(channel int, secondId string) *types.Account
	AccountByUniqueCode(uniqueCode string) *types.Account
}

//db
type backend struct {
}

func (b *backend) AccountByTypeAndSecondId(at int, secondId string) (acc *types.Account) {
	return
}

func (b *backend) AccountByUniqueCode(uniqueCode string) (acc *types.Account) {
	return
}

func NewBackend() Backend {
	b := &backend{}
	return b
}
