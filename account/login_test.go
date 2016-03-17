package account_test

import (
	"bytes"

	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
	"github.com/xozrc/account/account"
	"github.com/xozrc/account/pkg/stringutil"
	"github.com/xozrc/account/types"
)

type testBackend struct {
	store []*types.Account
}

func (b *testBackend) Login(channel int, secondId string, uniqueCode string) (*types.Account, error) {
	acc := b.accountByChannelAndSecondId(channel, secondId)
	if acc == nil {
		acc = types.NewAccount(1, secondId, channel, uniqueCode, "")
	}
	return acc, nil
}

func (b *testBackend) LoginUniqueCode(uniqueCode string) (*types.Account, error) {

	//unique code valid
	if !stringutil.IsUniqueCodeValid(uniqueCode) {
		return nil, account.UniqueCodeInvalidError
	}

	acc := b.accountByUniqueCode(uniqueCode)
	if acc == nil {
		acc = types.NewAccount(1, "", int(types.Visitor), uniqueCode, "")
	}

	return acc, nil
}

func (b *testBackend) accountByUniqueCode(uniqueCode string) *types.Account {
	for _, acc := range b.store {
		if strings.EqualFold(uniqueCode, acc.UniqueCode) {
			return acc
		}
	}
	return nil
}

func (b *testBackend) accountByChannelAndSecondId(channel int, secondId string) *types.Account {
	for _, acc := range b.store {
		if channel == acc.Channel && strings.EqualFold(acc.SecondId, secondId) {
			return acc
		}
	}
	return nil
}

func NewBackend(s []*types.Account) account.Backend {
	return &testBackend{store: s}
}

func TestLogin(t *testing.T) {
	//todo :mock db and redis

	ts := httptest.NewServer(http.HandlerFunc(account.Login))
	defer ts.Close()

	content := bytes.NewBufferString("")
	_, err := http.Post(ts.URL, "application/json", content)
	assert.NoError(t, err, "login error")

}
