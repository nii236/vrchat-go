package client

import "errors"

type ErrorResponse struct {
	Err struct {
		Message    string `json:"message"`
		StatusCode int    `json:"status_code"`
	} `json:"error"`
}

func (e *ErrorResponse) Unwrap() error {
	return errors.New(e.Err.Message)
}
func (e *ErrorResponse) Error() string {
	return e.Err.Message
}
