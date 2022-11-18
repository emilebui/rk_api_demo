package wagerlogic

import (
	"github.com/stretchr/testify/assert"
	"rk_echo/db/models"
	"rk_echo/db/sqldb"
	wagerdto "rk_echo/service/wager/dto"
	"testing"
)

func TestGet(t *testing.T) {

	// Init Fake DB
	sqldb.InitFakeDB()

	// Test when there is no wager in db, get should return error
	_, err := Get(1)
	if err == nil {
		t.Fatal("This should fail")
	}

	// Test when there is wager in the db, get should return result

	//Init Data
	sqldb.DB.Create(&models.Wager{Id: 1})
	_, err = Get(1)
	if err != nil {
		t.Fatal("This should not fail")
	}
}

func TestCreateTransaction(t *testing.T) {

	// Init DB
	sqldb.InitFakeDB()

	// Test Create without valid
	input := &wagerdto.BuyWagerInput{
		WagerId:     1,
		BuyingPrice: 200,
	}

	_, err := CreateTransaction(input)
	if err == nil {
		t.Fatal("This should fail")
	}

	// Test Create with valid wagerId
	sqldb.DB.Create(&models.Wager{Id: 1})
	_, err = CreateTransaction(input)
	if err != nil {
		t.Fatal("This should not fail")
	}
}

func TestBuy(t *testing.T) {

	// Init Fake DB
	sqldb.InitFakeDB()

	// input
	input := &wagerdto.BuyWagerInput{
		WagerId:     1,
		BuyingPrice: 200,
	}

	// Buy with wrong wager id --> fail
	_, err := Buy(input)
	if err == nil {
		t.Fatal("This should fail")
	}

	//Init wager
	sqldb.DB.Create(&models.Wager{Id: 1, SellingPrice: 100, CurrentSellingPrice: 100})

	// Buy with buying price > current selling price --> fail
	_, err = Buy(input)
	if err == nil {
		t.Fatal("This should fail")
	}

	// Change wager to lower the buying price
	input.BuyingPrice = 100

	// Buy with the correct buying price --> pass
	_, err = Buy(input)
	if err != nil {
		t.Fatal("this should not fail")
	}

	// check result

	wager, err := Get(1)
	if err != nil {
		t.Fatal("This should not fail")
	}

	assert.Equal(t, 0.0, wager.CurrentSellingPrice)
}
