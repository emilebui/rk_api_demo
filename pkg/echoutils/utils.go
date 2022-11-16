package echoutils

import (
	"github.com/labstack/echo/v4"
	"rk_echo/pkg/errutil"
	"rk_echo/pkg/validate"
)

type ResponseMessage struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func EchoGetInputFromCtx(ctx echo.Context, obj interface{}) errutil.EchoError {
	if err := ctx.Bind(obj); err != nil {
		return &errutil.EchoErrorStruct{
			Code:  500,
			Error: err.Error(),
		}
	}
	return nil
}

func EchoGetInputAndValidate(ctx echo.Context, obj interface{}) errutil.EchoError {
	err := EchoGetInputFromCtx(ctx, obj)
	if err != nil {
		return err
	}

	errs := validate.Validate(obj)
	if errs != nil {
		return &errutil.EchoErrorStruct{
			Code:  422,
			Error: "Input is not valid",
			Info:  errs,
		}
	}
	return nil
}

func ReturnDBResult(err error) errutil.EchoError {
	if err != nil {
		return &errutil.EchoErrorStruct{
			Code:  400,
			Error: err.Error(),
		}
	}
	return nil
}
