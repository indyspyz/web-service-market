package entity

type Market struct {
	Id         int    `json:"id"`
	MarketName string `json:"marketname"`
}

type Product struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	MarketId int    `json:"marketid"`
}
