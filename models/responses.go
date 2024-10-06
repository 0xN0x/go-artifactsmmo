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

type ItemReponse struct {
	Cooldown  Cooldown   `json:"cooldown"`
	Item      SimpleItem `json:"item"`
	Character Character  `json:"character"`
}

type CharacterFight struct {
	Cooldown  Cooldown  `json:"cooldown"`
	Fight     Fight     `json:"fight"`
	Character Character `json:"character"`
}

type BankExtensionTransaction struct {
	Cooldown    Cooldown      `json:"cooldown"`
	Transaction BankExtension `json:"transaction"`
	Character   Character     `json:"character"`
}

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

type TaskData struct {
	Cooldown  Cooldown  `json:"cooldown"`
	Task      Task      `json:"task"`
	Character Character `json:"character"`
}

type TaskTradeData struct {
	Cooldown  Cooldown  `json:"cooldown"`
	Trade     TaskTrade `json:"trade"`
	Character Character `json:"character"`
}

type TaskCancelled struct {
	Cooldown  Cooldown  `json:"cooldown"`
	Character Character `json:"character"`
}

type TaskRewardData struct {
	Cooldown  Cooldown   `json:"cooldown"`
	Reward    TaskReward `json:"reward"`
	Character Character  `json:"character"`
}
