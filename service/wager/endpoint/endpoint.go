package wagerep

import (
	"github.com/labstack/echo/v4"
	rkecho "github.com/rookie-ninja/rk-echo/boot"
	"net/http"
	"rk_echo/pkg/dbutils"
	"rk_echo/pkg/echoutils"
	wagerdto "rk_echo/service/wager/dto"
	wagerlogic "rk_echo/service/wager/logic"
)

func RegisterAPI(echoEntry *rkecho.EchoEntry) {
	echoEntry.Echo.POST("/wagers", PlaceWager)
	echoEntry.Echo.POST("/buy/:wager_id", BuyWager)
	echoEntry.Echo.GET("/wagers", ListWager)
}

// PlaceWager handler
// @Summary Place a wager
// @version 1.0
// @produce application/json
// @Tags Wager
// @Param input body wagerdto.PlaceWagerInput true "Input Required"
// @Success 201 {object} models.Wager
// @Failure default {object} errutil.EchoErrorStruct
// @Router /wagers [post]
func PlaceWager(ctx echo.Context) error {

	input := new(wagerdto.PlaceWagerInput)

	err := echoutils.EchoGetInputAndValidate(ctx, input)
	if err != nil {
		return err.ReturnErrorMessage(ctx)
	}

	wager, err := wagerlogic.Place(input)
	if err != nil {
		return err.ReturnErrorMessage(ctx)
	}

	return ctx.JSON(http.StatusCreated, wagerdto.Struct2PlaceWagerResponse(wager))
}

// BuyWager handler
// @Summary Buy a wager
// @version 1.0
// @produce application/json
// @Tags Wager
// @Param wager_id path string true "ID of the wager to buy"
// @Param input body wagerdto.BuyWagerInput true "Input Required"
// @Success 201 {object} models.Transaction
// @Failure default {object} errutil.EchoErrorStruct
// @Router /buy/{wager_id} [post]
func BuyWager(ctx echo.Context) error {

	input := new(wagerdto.BuyWagerInput)

	err := echoutils.EchoGetInputAndValidate(ctx, input)
	if err != nil {
		return err.ReturnErrorMessage(ctx)
	}

	id, err := echoutils.String2Uint(ctx.Param("wager_id"))
	if err != nil {
		return err.ReturnErrorMessage(ctx)
	}
	input.WagerId = id

	transaction, err := wagerlogic.Buy(input)
	if err != nil {
		return err.ReturnErrorMessage(ctx)
	}

	return ctx.JSON(http.StatusCreated, wagerdto.Struct2BuyWagerResponse(transaction))
}

// ListWager handler
// @Summary List all wager
// @version 1.0
// @produce application/json
// @Tags Wager
// @Param page query string false "Page"
// @Param limit query string false "Limit"
// @Param sort query string false "Sort"
// @Success 201 {object} models.Transaction
// @Failure default {object} errutil.EchoErrorStruct
// @Router /wagers [get]
func ListWager(ctx echo.Context) error {

	page, err := echoutils.String2Int(ctx.QueryParam("page"))
	if err != nil {
		return err.ReturnErrorMessage(ctx)
	}
	limit, err := echoutils.String2Int(ctx.QueryParam("limit"))
	if err != nil {
		return err.ReturnErrorMessage(ctx)
	}
	sort := ctx.QueryParam("sort")

	pagingConf := dbutils.PaginateInput{
		Page: page,
		Size: limit,
		Sort: sort,
	}
	wagerList, err := wagerlogic.List(pagingConf)
	if err != nil {
		return err.ReturnErrorMessage(ctx)
	}
	return ctx.JSON(http.StatusOK, wagerdto.ToListResponse(wagerList))
}
