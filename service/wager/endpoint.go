package wager

import (
	"github.com/labstack/echo/v4"
	rkecho "github.com/rookie-ninja/rk-echo/boot"
	"net/http"
	"rk_echo/pkg/echoutils"
)

func RegisterAPI(echoEntry *rkecho.EchoEntry) {
	echoEntry.Echo.POST("/wagers", PlaceWagerEp)
}

type PlaceWagerInput struct {
	TotalWagerValue   int64 `json:"total_wager_value" validate:"required,gt=0"`
	Odd               int32 `json:"odd" validate:"required,gt=0"`
	SellingPercentage int32 `json:"selling_percentage" validate:"required,gte=1,lte=100"`
	SellingPrice      int32 `json:"selling_price" validate:"required"`
}

// PlaceWagerEp handler
// @Summary Place a wager
// @version 1.0
// @produce application/json
// @Tags Wager
// @Param input body PlaceWagerInput true "Input Required"
// @Success 201 {object} models.Wager
// @Failure default {object} errutil.EchoErrorStruct
// @Router /wagers [post]
func PlaceWagerEp(ctx echo.Context) error {

	input := new(PlaceWagerInput)

	err := echoutils.EchoGetInputAndValidate(ctx, input)
	if err != nil {
		return err.ReturnErrorMessage(ctx)
	}

	return ctx.String(http.StatusCreated, "OK!")
}
