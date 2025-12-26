package model

import (
	"gorm.io/gorm"
)

type CreateAuctionEvent struct {
	gorm.Model
	StartPrice int64
	AuctionId  int64
	NftAddress string
	TokenId    int64
}
