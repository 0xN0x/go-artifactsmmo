package models

type GETransaction struct {
	Code       int `json:"code"`
	Quantity   int `json:"quantity"`
	Price      int `json:"price"`
	TotalPrice int `json:"total_price"`
}

type GEItem struct {
	Code     string `json:"code"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
}

type GEItemSchema struct {
	Code        string `json:"code"`
	Stock       int    `json:"stock"`
	SellPrice   int    `json:"sell_price"`
	BuyPrice    int    `json:"buy_price"`
	MaxQuantity int    `json:"max_quantity"`
}
