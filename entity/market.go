package entity

type Market struct {
	Id         int    `json:"id"`
	MarketName string `json:"marketname"`
}

type Product struct {
	Id       int    `json:"id"`
	Name     string `json:"productname"`
	Price    int    `json:"productprice"`
	MarketId int    `json:"marketid"`
}
