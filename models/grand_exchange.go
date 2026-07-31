package models

type GEType string

const (
	Sell   GEType = "sell"
	Buy    GEType = "buy"
	GENone GEType = ""
)

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

type GEOrderSchema struct {
	Id        string `json:"id"`
	Type      GEType `json:"type"`
	Account   string `json:"account"`
	Code      string `json:"code"`
	Quantity  int    `json:"quantity"`
	Price     int    `json:"price"`
	CreatedAt string `json:"created_at"`
}
}
