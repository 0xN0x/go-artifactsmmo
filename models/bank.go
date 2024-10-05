package models

type BankItemTransaction struct {
	Cooldown  Cooldown     `json:"cooldown"`
	Item      Item         `json:"item"`
	Bank      []SimpleItem `json:"bank"`
	Character Character    `json:"character"`
}

type BankGoldTransaction struct {
	Cooldown  Cooldown  `json:"cooldown"`
	Bank      Gold      `json:"bank"`
	Character Character `json:"character"`
}

type BankExtension struct {
	Price int `json:"price"`
}

type BankExtensionTransaction struct {
	Cooldown    Cooldown      `json:"cooldown"`
	Transaction BankExtension `json:"transaction"`
	Character   Character     `json:"character"`
}
