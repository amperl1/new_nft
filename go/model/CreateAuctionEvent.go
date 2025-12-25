package model

import (
	"gorm.io/gorm"
)

type CreateAuctionEvent struct {
	gorm.Model
	StartPrice string
	AuctionId  string
	NftAddress string
	TokenId    string
}
