package account_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

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
)

var _ = fmt.Print

func TestLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mb := mock_account.NewMockAccountBackend(ctrl)

	ctx := context.TODO()
	cctx := account.WithAccountBackend(ctx, mb)

	ts := httptest.NewServer(account.LoginHTTPHandler(cctx))
	defer ts.Close()

	//visitor account
	testVisitor := types.NewAccount(1, "", int(types.Visitor), "visitor", "visitor_name")
	mb.EXPECT().Login(types.AccountType(testVisitor.AccountType), testVisitor.SecondId, testVisitor.UniqueCode).Return(testVisitor, nil)

	//	content := bytes.NewBufferString("")
	form := &account.LoginForm{}
	form.AccountType = testVisitor.AccountType
	form.SecondId = testVisitor.SecondId
	form.UniqueCode = testVisitor.UniqueCode
	content, _ := json.Marshal(form)

	_, err := http.Post(ts.URL, "application/json", bytes.NewBuffer(content))
	assert.NoError(t, err, "login error")
}

func TestLoginBindError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.TODO()

	ts := httptest.NewServer(account.LoginHTTPHandler(ctx))
	defer ts.Close()

	res, err := http.Post(ts.URL, "application/json", bytes.NewBufferString("error"))
	assert.NoError(t, err, "login error")
	assert.True(t, res.StatusCode == http.StatusBadRequest, "error response status")
}
