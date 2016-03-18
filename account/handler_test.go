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
	"golang.org/x/net/context"
	"github.com/xozrc/rest"
)

type testBackend struct {
	store []*types.Account
}

func (b *testBackend) AccountByTypeAndSecondId(channel int,secondId string)*types.Account {
	for _, acc := range b.store {
		if channel == acc.Channel && strings.EqualFold(acc.SecondId, secondId) {
			return acc
		}
	}
	return nil
}

func (b *testBackend) AccountByUniqueCode(uniqueCode string) *types.Account){

	for _, acc := range b.store {
		if strings.EqualFold(uniqueCode, acc.UniqueCode) {
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
	ctx := context.TODO()
	as := account.NewAccountService()

	cctx := account.WithAccountService(ctx,as)
	ts := httptest.NewServer(rest.RestHandler(cctx, rest.HandleFunc(Login))
	defer ts.Close()

	content := bytes.NewBufferString("")
	_, err := http.Post(ts.URL, "application/json", content)
	assert.NoError(t, err, "login error")

}
