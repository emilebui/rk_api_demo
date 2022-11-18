package wagerdto

import (
	"math"
	"rk_echo/db/models"
	"rk_echo/pkg/helpers"
)

type PlaceWagerInput struct {
	TotalWagerValue   float64 `json:"total_wager_value" validate:"required,gt=0"`
	Odds              float64 `json:"odds" validate:"required,gt=0"`
	SellingPercentage float64 `json:"selling_percentage" validate:"required,gte=1,lte=100"`
	SellingPrice      float64 `json:"selling_price" validate:"required"`
}

type BuyWagerInput struct {
	WagerId     uint    `json:"-"`
	BuyingPrice float64 `json:"buying_price" validate:"required,gt=0"`
}

func Struct2PlaceWagerResponse(wager *models.Wager) *models.Wager {
	wager.Odds = math.Round(wager.Odds)
	wager.TotalWagerValue = math.Round(wager.TotalWagerValue)
	wager.SellingPercentage = math.Round(wager.SellingPercentage)
	wager.SellingPrice = helpers.RoundFloat(wager.SellingPrice)
	wager.CurrentSellingPrice = helpers.RoundFloat(wager.CurrentSellingPrice)
	wager.PercentageSold = math.Round(wager.PercentageSold)
	wager.AmountSold = helpers.RoundFloat(wager.AmountSold)
	return wager
}

func Struct2BuyWagerResponse(transaction *models.Transaction) *models.Transaction {
	transaction.BuyingPrice = math.Round(transaction.BuyingPrice)
	return transaction
}

func ToListResponse(input *[]models.Wager) *[]models.Wager {
	arr := *input
	newArr := make([]models.Wager, len(arr))
	for i, s := range arr {
		newArr[i] = *Struct2PlaceWagerResponse(&s)
	}
	return &newArr
}
