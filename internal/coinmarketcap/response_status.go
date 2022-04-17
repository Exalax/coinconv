package coinmarketcap

import "fmt"

// statusResponse is a response status data.
type statusResponse struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

// Err returns error if status is an error status, nil otherwise.
func (s statusResponse) Err() error {
	if s.ErrorCode > 0 {
		return fmt.Errorf("%w %d: %s", ErrStatusError, s.ErrorCode, s.ErrorMessage)
	}
	return nil
}
