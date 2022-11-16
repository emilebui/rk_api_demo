package demo

import (
	"github.com/labstack/echo/v4"
	rkecho "github.com/rookie-ninja/rk-echo/boot"
	"net/http"
)

func RegisterAPI(echoEntry *rkecho.EchoEntry) {
	echoEntry.Echo.GET("/v1/greeter", HelloWorld)
}

var logic Interface = &Logic{}

// HelloWorld handler
// @Summary Say hello to the current user
// @version 1.0
// @produce application/json
// @Tags Demo
// @Success 200 {object} GreeterResponse
// @Failure default {object} errutil.EchoErrorStruct
// @Router /v1/greeter [get]
func HelloWorld(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, logic.Greeter())
}
