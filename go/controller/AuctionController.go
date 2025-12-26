package controller

import (
	"new_nft_go/database"
	"new_nft_go/model"
	"new_nft_go/model/response"

	"github.com/gin-gonic/gin"
)

type AuctionController struct{}

type GetCreateAuctionsRequest struct {
	AuctionId string `json:"auctionId"`
}

func (ac *AuctionController) GetCreateAuctions(c *gin.Context) {
	var req GetCreateAuctionsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.StatusUnauthorized(c, "params error!")
		return
	}
	AuctionId := req.AuctionId
	if AuctionId == "" {
		var createAuctionEvents []model.CreateAuctionEvent
		database.Db.Find(&createAuctionEvents)
		response.Success(c, createAuctionEvents)
	} else {
		var createAuctionEvent []model.CreateAuctionEvent
		database.Db.Find(&createAuctionEvent, "auction_id = ?", c.MustGet("auction_id"))
		response.Success(c, createAuctionEvent)
	}
}

type GetBidAuctionsRequest struct {
	AuctionId string `json:"auctionId"`
}

func (ac *AuctionController) GetBidAuctions(c *gin.Context) {
	var req GetBidAuctionsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.StatusUnauthorized(c, "params error!")
		return
	}
	AuctionId := req.AuctionId
	if AuctionId == "" {
		var bidAuctionEvents []model.BidAuctionEvent
		database.Db.Find(&bidAuctionEvents)
		response.Success(c, bidAuctionEvents)
	} else {
		var bidAuctionEvent []model.BidAuctionEvent
		database.Db.Find(&bidAuctionEvent, "auction_id = ?", c.MustGet("auction_id"))
		response.Success(c, bidAuctionEvent)
	}
}

type GetEndAuctionsRequest struct {
	AuctionId string `json:"auctionId"`
}

func (ac *AuctionController) GetEndAuctions(c *gin.Context) {
	var req GetEndAuctionsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.StatusUnauthorized(c, "params error!")
		return
	}
	AuctionId := req.AuctionId
	if AuctionId == "" {
		var endAuctionEvents []model.EndAuctionEvent
		database.Db.Find(&endAuctionEvents)
		response.Success(c, endAuctionEvents)
	} else {
		var endAuctionEvent []model.EndAuctionEvent
		database.Db.Find(&endAuctionEvent, "auction_id = ?", c.MustGet("auction_id"))
		response.Success(c, endAuctionEvent)
	}
}
