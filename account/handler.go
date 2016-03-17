package account

import (
	"fmt"
	"net/http"
)

var _ = fmt.Print

func Login(b Backend) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(rw, "hello")
		// tId := 1
		// account, err := b.Login(tId)
		// if err != nil {
		// 	return
		// }

	}
}
