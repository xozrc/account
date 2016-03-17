package account

import (
	"github.com/xozrc/account/types"
)

type Backend interface {
	Login(channel int, secondId string, uniqueCode string) (*types.Account, error)
	LoginUniqueCode(uniqueCode string) (*types.Account, error)
}

type backend struct {
}

func (b *backend) Login(channel int, id string, uniqueCode string) (acc *types.Account, err error) {
	return
}

func (b *backend) LoginUniqueCode(uniqueCode string) (acc *types.Account, err error) {
	return
}

func NewBackend() Backend {
	b := &backend{}
	return b
}
