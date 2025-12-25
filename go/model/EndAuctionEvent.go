package model

import (
	"gorm.io/gorm"
)

type EndAuctionEvent struct {
	gorm.Model
	AuctionId string
}
