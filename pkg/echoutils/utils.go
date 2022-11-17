package echoutils

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"rk_echo/pkg/errutil"
	"rk_echo/pkg/validate"
	"strconv"
)

type ResponseMessage struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func EchoGetInputFromCtx(ctx echo.Context, obj interface{}) errutil.EchoError {
	if err := ctx.Bind(obj); err != nil {
		return errutil.NewEchoError(500, err.Error())
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
		return errutil.NewEchoError(422, "Input is not valid", errs)
	}
	return nil
}

func ReturnDBResult(err error) errutil.EchoError {
	if err != nil {
		return errutil.NewEchoError(400, err.Error())
	}
	return nil
}

func String2Uint(str string) (uint, errutil.EchoError) {
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0, errutil.NewEchoError(400, "This is not a valid Id")
	}

	return uint(i), nil
}

func String2Int(str string) (int, errutil.EchoError) {
	if str == "" {
		return 0, nil
	}
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0, errutil.NewEchoError(400, fmt.Sprintf("%s is not a valid integer", str))
	}
	return i, nil
}
