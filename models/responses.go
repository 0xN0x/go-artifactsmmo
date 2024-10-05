package models

type CharacterMovementData struct {
	Cooldown    Cooldown    `json:"cooldown"`
	Destination Destination `json:"destination"`
	Character   Character   `json:"character"`
}

type EquipRequest struct {
	Cooldown  Cooldown  `json:"cooldown"`
	Slot      Slot      `json:"slot"`
	Item      Item      `json:"item"`
	Character Character `json:"character"`
}

type SkillData struct {
	Cooldown  Cooldown  `json:"cooldown"`
	Details   Detail    `json:"details"`
	Character Character `json:"character"`
}

type Recycling struct {
	Cooldown  Cooldown   `json:"cooldown"`
	Details   ItemsArray `json:"details"`
	Character Character  `json:"character"`
}

type BankTransaction struct {
	Cooldown    Cooldown `json:"cooldown"`
	Transaction struct {
		Price int `json:"price"`
	} `json:"transaction"`
	Character Character `json:"character"`
}

type GETransactionResponse struct {
	Cooldown    Cooldown      `json:"cooldown"`
	Transaction GETransaction `json:"transaction"`
	Character   Character     `json:"character"`
}
