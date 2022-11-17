package wagerlogic

import (
	"rk_echo/db/models"
	"rk_echo/db/sqldb"
	"rk_echo/pkg/echoutils"
	"rk_echo/pkg/errutil"
	wagerdto "rk_echo/service/wager/dto"
)

func Place(input *wagerdto.PlaceWagerInput) (*models.Wager, errutil.EchoError) {

	if input.SellingPrice <= (input.TotalWagerValue * (input.SellingPercentage / 100)) {
		return nil, errutil.NewEchoError(400, "Selling price must be greater than `total_wager_value` * `selling_percentage`")
	}

	newWager := &models.Wager{
		TotalWagerValue:     input.TotalWagerValue,
		Odds:                input.Odds,
		SellingPercentage:   input.SellingPercentage,
		SellingPrice:        input.SellingPrice,
		CurrentSellingPrice: input.SellingPrice,
	}

	err := sqldb.DB.Create(newWager).Error
	return newWager, echoutils.ReturnDBResult(err)
}
