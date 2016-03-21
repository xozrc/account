package account_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	_ "strings"
	"testing"
)

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

import (
	"github.com/xozrc/account/account"
	"github.com/xozrc/account/account/mock_account"

	"github.com/xozrc/account/types"
	"github.com/xozrc/rest"
)

func TestLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mb := mock_account.NewMockAccountBackend(ctrl)

	ctx := context.TODO()
	cctx := account.WithAccountBackend(ctx, mb)

	ts := httptest.NewServer(rest.RestHandler(cctx, rest.HandleFunc(account.Login)))
	defer ts.Close()

	//visitor account
	testVisitor := types.NewAccount(1, "", int(types.Visitor), "visitor", "visitor_name")
	mb.EXPECT().Login(types.AccountType(testVisitor.AccountType), testVisitor.SecondId, testVisitor.UniqueCode).Return(testVisitor, nil)

	content := bytes.NewBufferString("")
	_, err := http.Post(ts.URL, "application/json", content)
	assert.NoError(t, err, "login error")

}
