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

type BaseAchievement struct {
	Name        string          `json:"name"`
	Code        string          `json:"code"`
	Description string          `json:"description"`
	Points      int             `json:"points"`
	Type        AchievementType `json:"type"`
	Target      string          `json:"target"`
	Total       int             `json:"total"`
}

type TaskFull struct {
	Code        string   `json:"code"`
	Level       int      `json:"level"`
	Type        TaskType `json:"type"`
	MinQuantity int      `json:"min_quantity"`
	MaxQuantity int      `json:"max_quantity"`
	Skill       string   `json:"skill"`
}

type TaskRewardFull struct {
	Code        string  `json:"code"`
	MinQuantity int     `json:"min_quantity"`
	MaxQuantity int     `json:"max_quantity"`
	Odds        float32 `json:"odds"`
}

type GEItems struct {
	Code        string `json:"code"`
	Stock       int    `json:"stock"`
	SellPrice   int    `json:"sell_price"`
	BuyPrice    int    `json:"buy_price"`
	MaxQuantity int    `json:"max_quantity"`
}

type Resource struct {
	Name  string     `json:"name"`
	Code  string     `json:"code"`
	Skill string     `json:"skill"`
	Level int        `json:"level"`
	Drops []DropFull `json:"drops"`
}

type Monster struct {
	Name        string     `json:"name"`
	Code        string     `json:"code"`
	Level       int        `json:"level"`
	Hp          int        `json:"hp"`
	AttackFire  int        `json:"attack_fire"`
	AttackEarth int        `json:"attack_earth"`
	AttackWater int        `json:"attack_water"`
	AttackAir   int        `json:"attack_air"`
	ResFire     int        `json:"res_fire"`
	ResEarth    int        `json:"res_earth"`
	ResWater    int        `json:"res_water"`
	ResAir      int        `json:"res_air"`
	MinGold     int        `json:"min_gold"`
	MaxGold     int        `json:"max_gold"`
	Drops       []DropFull `json:"drops"`
}

type SingleItem struct {
	Item Item         `json:"item"`
	GE   GEItemSchema `json:"ge"`
}

type MapSchema Destination
