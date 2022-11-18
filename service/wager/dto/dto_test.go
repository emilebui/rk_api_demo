package wagerdto

import (
	"github.com/stretchr/testify/assert"
	"rk_echo/db/models"
	"testing"
)

func TestStruct2PlaceWagerResponse(t *testing.T) {
	obj := &models.Wager{
		TotalWagerValue:     99.9999,
		Odds:                99.88,
		SellingPercentage:   68.9,
		SellingPrice:        200.23232,
		CurrentSellingPrice: 15.8232,
		PercentageSold:      22.2,
		AmountSold:          172.289,
	}
	obj = Struct2PlaceWagerResponse(obj)

	assert.Equal(t, 100, obj.TotalWagerValue)
	assert.Equal(t, 100, obj.Odds)
	assert.Equal(t, 69, obj.SellingPercentage)
	assert.Equal(t, 200.23, obj.SellingPrice)
	assert.Equal(t, 15.82, obj.CurrentSellingPrice)
	assert.Equal(t, 22, obj.PercentageSold)
	assert.Equal(t, 172.29, obj.AmountSold)
}

func TestStruct2BuyWagerResponse(t *testing.T) {

	// Case 1

	obj := &models.Transaction{
		BuyingPrice: 200.99,
	}
	obj = Struct2BuyWagerResponse(obj)
	assert.Equal(t, 200.99, obj.BuyingPrice)

	// Case 2

	obj = &models.Transaction{
		BuyingPrice: 200.369,
	}
	obj = Struct2BuyWagerResponse(obj)
	assert.Equal(t, 200.37, obj.BuyingPrice)

}

func TestToListResponse(t *testing.T) {

	// Init Data
	arr := []models.Wager{
		{
			TotalWagerValue:     99.9999,
			Odds:                99.88,
			SellingPercentage:   68.9,
			SellingPrice:        200.23232,
			CurrentSellingPrice: 15.8232,
			PercentageSold:      22.2,
			AmountSold:          172.289,
		},
		{
			TotalWagerValue:     99.9999,
			Odds:                99.88,
			SellingPercentage:   68.9,
			SellingPrice:        200.23232,
			CurrentSellingPrice: 15.8232,
			PercentageSold:      22.2,
			AmountSold:          172.289,
		},
	}

	temp := ToListResponse(&arr)

	for _, s := range *temp {
		assert.Equal(t, 100, s.TotalWagerValue)
		assert.Equal(t, 100, s.Odds)
		assert.Equal(t, 69, s.SellingPercentage)
		assert.Equal(t, 200.23, s.SellingPrice)
		assert.Equal(t, 15.82, s.CurrentSellingPrice)
		assert.Equal(t, 22, s.PercentageSold)
		assert.Equal(t, 172.29, s.AmountSold)
	}
}
