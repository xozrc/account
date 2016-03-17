package router_test

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/xozrc/account/router"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogin(t *testing.T) {
	//todo :mock db and redis

	// ts := httptest.NewServer(http.HandlerFunc(router.Login))
	// defer ts.Close()

	// content := bytes.NewBufferString("")
	// _, err := http.Post(ts.URL, "application/json", content)
	// assert.NoError(t, err, "login error")

}
