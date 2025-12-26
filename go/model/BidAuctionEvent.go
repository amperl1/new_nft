package model

import (
	"gorm.io/gorm"
)

type BidAuctionEvent struct {
	gorm.Model
	AuctionId    int64
	Amount       int64
	TokenAddress string
}
