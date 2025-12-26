package server

import (
	"fmt"
	"new_nft_go/controller"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	auctionController := &controller.AuctionController{}
	v1 := r.Group("/api/v1")
	{
		v1.GET("/getCreateAuctions", auctionController.GetCreateAuctions)
		v1.GET("/getEndAuctions", auctionController.GetEndAuctions)
		v1.GET("/getBidAuctions", auctionController.GetBidAuctions)
	}
	fmt.Println("Starting server")
	return r
}
