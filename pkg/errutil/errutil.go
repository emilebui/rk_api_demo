package errutil

import "github.com/labstack/echo/v4"

type EchoError interface {
	ReturnErrorMessage(ctx echo.Context) error
}

type EchoErrorStruct struct {
	Error string      `json:"error"`
	Code  int         `json:"code"`
	Info  interface{} `json:"info,omitempty"`
}

func NewEchoError(code int, msg string, info ...interface{}) EchoError {

	var i interface{}
	if len(info) > 0 {
		i = info[0]
	}

	return &EchoErrorStruct{
		Error: msg,
		Code:  code,
		Info:  i,
	}
}

func (e *EchoErrorStruct) ReturnErrorMessage(ctx echo.Context) error {
	return ctx.JSON(e.Code, e)
}

func PanicHandler(err error) {
	if err != nil {
		panic(err)
	}
}
