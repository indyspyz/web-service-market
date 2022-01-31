package entity

type Market struct {
	MarketName string `json:"marketname"`
}

type Product struct {
	Name  string `json:"name"`
	Price int32  `json:"price"`
}
