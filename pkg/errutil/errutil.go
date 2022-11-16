package errutil

import "github.com/labstack/echo/v4"

type EchoError interface {
	ReturnErrorMessage(ctx echo.Context) error
}

type EchoErrorStruct struct {
	Error string      `json:"error"`
	Code  int         `json:"code"`
	Info  interface{} `json:"info"`
}

func (e *EchoErrorStruct) ReturnErrorMessage(ctx echo.Context) error {
	return ctx.JSON(e.Code, e)
}

func PanicHandler(err error) {
	if err != nil {
		panic(err)
	}
}
