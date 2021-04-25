package errcode

import "fmt"

type MyError struct {
	Status  int         `json:"-"`
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (e *MyError) Error() string {
	return fmt.Sprintf("error: status = %d code= %s desc = %s data = %+v", e.Status, e.Code, e.Message, e.Data)
}
