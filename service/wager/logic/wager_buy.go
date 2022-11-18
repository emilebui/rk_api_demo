package wagerlogic

import (
	"rk_echo/db/models"
	"rk_echo/db/sqldb"
	"rk_echo/pkg/echoutils"
	"rk_echo/pkg/errutil"
	wagerdto "rk_echo/service/wager/dto"
)

func Get(id uint) (*models.Wager, errutil.EchoError) {
	var wager models.Wager
	err := sqldb.DB.First(&wager, "id = ?", id).Error
	return &wager, echoutils.ReturnDBResult(err)
}

func CreateTransaction(input *wagerdto.BuyWagerInput) (*models.Transaction, errutil.EchoError) {
	newTransaction := &models.Transaction{
		WagerId:     input.WagerId,
		BuyingPrice: input.BuyingPrice,
	}
	err := sqldb.DB.Create(newTransaction).Error
	return newTransaction, echoutils.ReturnDBResult(err)
}

func Buy(input *wagerdto.BuyWagerInput) (*models.Transaction, errutil.EchoError) {

	// Get Wager
	wager, err := Get(input.WagerId)
	if err != nil {
		return nil, err
	}

	// Check BuyingPrice <= CurrentSellingPrice
	if input.BuyingPrice > wager.CurrentSellingPrice {
		return nil, errutil.NewEchoError(400, "Buying price should not be greater than current selling price")
	}

	// Create Transaction
	transaction, err := CreateTransaction(input)
	if err != nil {
		return nil, err
	}

	// Update CurrentSellingPrice
	wager.CurrentSellingPrice = wager.CurrentSellingPrice - input.BuyingPrice

	// Update AmountSold
	wager.AmountSold += input.BuyingPrice

	// Update PercentageSold
	wager.PercentageSold = (wager.AmountSold / wager.SellingPrice) * 100

	// Save
	dbErr := sqldb.DB.Model(&wager).Select("*").Updates(&wager).Error
	return transaction, echoutils.ReturnDBResult(dbErr)
}
