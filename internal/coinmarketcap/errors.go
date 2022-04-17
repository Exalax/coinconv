package coinmarketcap

import (
	"errors"
)

// Errors.
var (
	ErrQueryConstructionFailed = errors.New("query construction failed")
	ErrQueryFailed             = errors.New("query failed")
	ErrResponseUnmarshalFailed = errors.New("response unmarshal failed")
	ErrStatusError             = errors.New("status error")
	ErrSymbolNotFound          = errors.New("response has no requested symbol")
)
