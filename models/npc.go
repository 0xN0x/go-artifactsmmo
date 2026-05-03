package models

type NPCTransaction struct {
	Code       string `json:"code"`
	Quantity   int    `json:"quantity"`
	Price      int    `json:"price"`
	TotalPrice int    `json:"total_price"`
}

type NPCType string

const (
	NPCMerchant NPCType = "merchant"
)

type NPC struct {
	Name        string  `json:"name"`
	Code        string  `json:"code"`
	Description string  `json:"description"`
	Type        NPCType `json:"type"`
}

type NPCItem struct {
	Code      string `json:"code"`
	NPC       string `json:"npc"`
	BuyPrice  int    `json:"buy_price"`
	SellPrice int    `json:"sell_price"`
}
