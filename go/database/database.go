package database

import (
	"log"
	"new_nft_go/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDataBase() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	Db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// 自动迁移数据库表结构
	err := Db.AutoMigrate(
		&model.BidAuctionEvent{},
		&model.CreateAuctionEvent{},
		&model.EndAuctionEvent{},
	)
	if err != nil {
	}
	log.Println("Connected to database")
}
