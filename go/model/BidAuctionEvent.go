package model

import (
	"gorm.io/gorm"
)

type BidAuctionEvent struct {
	gorm.Model
	AuctionId    string
	Amount       string
	TokenAddress string
}
