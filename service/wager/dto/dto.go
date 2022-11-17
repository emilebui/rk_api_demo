package wagerdto

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
