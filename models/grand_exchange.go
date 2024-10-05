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
