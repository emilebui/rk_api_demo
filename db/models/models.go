package models

import (
	"gorm.io/gorm"
	"time"
)

// *************************************
// *************** Model ***************
// *************************************

type Base struct {
	CreatedAt time.Time      `yaml:"-" json:"-"`
	UpdatedAt time.Time      `yaml:"-" json:"-"`
	DeletedAt gorm.DeletedAt `yaml:"-" json:"-" gorm:"index"`
}

type Wager struct {
	Base
	Id                  uint          `yaml:"id" json:"id" gorm:"primaryKey;autoIncrement"`
	TotalWagerValue     float64       `yaml:"total_wager_value" json:"total_wager_value"`
	Odds                float64       `yaml:"odds" json:"odds"`
	SellingPercentage   float64       `yaml:"selling_percentage" json:"selling_percentage"`
	SellingPrice        float64       `yaml:"selling_price" json:"selling_price"`
	CurrentSellingPrice float64       `yaml:"current_selling_price" json:"current_selling_price"`
	PercentageSold      float64       `yaml:"percentage_sold" json:"percentage_sold"`
	AmountSold          float64       `yaml:"amount_sold" json:"amount_sold"`
	PlacedAt            time.Time     `yaml:"placed_at" json:"placed_at" gorm:"autoCreateTime"`
	Transactions        []Transaction `gorm:"constraint:OnDelete:CASCADE;foreignkey:WagerId;references:Id" json:"-"`
}

type Transaction struct {
	Base
	Id          uint      `yaml:"id" json:"id" gorm:"primaryKey;autoIncrement"`
	WagerId     uint      `yaml:"wager_id" json:"wager_id" gorm:"primaryKey;autoIncrement"`
	BuyingPrice float64   `yaml:"buying_price" json:"buying_price"`
	BoughtAt    time.Time `yaml:"bought_at" json:"bought_at" gorm:"autoCreateTime"`
}
