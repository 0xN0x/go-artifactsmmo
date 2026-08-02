package models

type GEType string

const (
	Sell   GEType = "sell"
	Buy    GEType = "buy"
	GENone GEType = ""
)

type GETransaction struct {
	Id         string `json;"id"`
	Code       string `json:"code"`
	Quantity   int    `json:"quantity"`
	Price      int    `json:"price"`
	TotalPrice int    `json:"total_price"`
}

type GEItem struct {
	Code     string `json:"code"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
}

type GEBuyItem struct {
	Id       string `json:"id"`
	Quantity int    `json:"quantity"`
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

type GEHistorySchema struct {
	Id       string `json:"order_id"`
	Seller   string `json:"seller"`
	Buyer    string `json:"buyer"`
	Code     string `json:"code"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
	SoldAt   string `json:"sold_at"`
}
