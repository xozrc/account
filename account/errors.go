package account

import (
	"github.com/xozrc/pkg/restutil"
)

const (
	UniqueCodeInvalid = 10000
)

var (
	UniqueCodeInvalidError = restutil.NewRestError(UniqueCodeInvalid, "username invalid")
)
