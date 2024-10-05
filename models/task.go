package models

type TaskType string

const (
	TaskTypeMonsters TaskType = "monsters"
	TaskTypeItems    TaskType = "items"
)

type Task struct {
	Code  string   `json:"code"`
	Type  TaskType `json:"type"`
	Total int      `json:"total"`
}

type TaskData struct {
	Cooldown  Cooldown  `json:"cooldown"`
	Task      Task      `json:"task"`
	Character Character `json:"character"`
}

type TaskReward struct {
	Code     string `json:"code"`
	Quantity int    `json:"quantity"`
}

type TaskRewardData struct {
	Cooldown  Cooldown   `json:"cooldown"`
	Reward    TaskReward `json:"reward"`
	Character Character  `json:"character"`
}

type TaskTrade SimpleItem

type TaskTradeData struct {
	Cooldown  Cooldown  `json:"cooldown"`
	Trade     TaskTrade `json:"trade"`
	Character Character `json:"character"`
}

type TaskCancelled struct {
	Cooldown  Cooldown  `json:"cooldown"`
	Character Character `json:"character"`
}
