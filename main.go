package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Cryptocurrency struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
	Rank   int    `json:"rank"`
	Image  string `json:"image"`
}

func main() {
	r := gin.Default()
	r.GET("/posts", func(c *gin.Context) {
		var cryptocurrencies = []Cryptocurrency{
			{Id: 1, Name: "Bitcoin", Symbol: "BTC", Rank: 1, Image: "https://s2.coinmarketcap.com/static/img/coins/64x64/1.png"},
			{Id: 2, Name: "Ethereum", Symbol: "ETH", Rank: 2, Image: "https://s2.coinmarketcap.com/static/img/coins/64x64/1027.png"},
			{Id: 3, Name: "Tether", Symbol: "USDT", Rank: 3, Image: "https://s2.coinmarketcap.com/static/img/coins/64x64/825.png"},
			{Id: 4, Name: "USD Coin", Symbol: "USDC", Rank: 4, Image: "https://s2.coinmarketcap.com/static/img/coins/64x64/3408.png"},
			{Id: 5, Name: "BNB", Symbol: "BNB", Rank: 5, Image: "https://s2.coinmarketcap.com/static/img/coins/64x64/1839.png"},
		}
		c.JSON(http.StatusOK, cryptocurrencies)
	})
	r.Run()
}
