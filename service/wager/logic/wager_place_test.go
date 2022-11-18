package wagerlogic

import (
	"github.com/stretchr/testify/assert"
	"rk_echo/db/sqldb"
	wagerdto "rk_echo/service/wager/dto"
	"testing"
)

func TestPlace(t *testing.T) {

	// Init DB
	sqldb.InitFakeDB()

	// Init a wager with selling price < total wager value * selling percentage
	input := &wagerdto.PlaceWagerInput{
		TotalWagerValue:   200,
		Odds:              100,
		SellingPercentage: 100,
		SellingPrice:      100,
	}

	_, err := Place(input)
	if err == nil {
		t.Fatal("this should fail")
	}

	// Changing to a valid wager
	input.SellingPrice = 300

	wager, err := Place(input)
	if err != nil {
		t.Fatal("this should not fail")
	}

	assert.Equal(t, 300.0, wager.SellingPrice)
	assert.Equal(t, 0.0, wager.AmountSold)
	assert.Equal(t, 0.0, wager.PercentageSold)
}
