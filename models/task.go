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

type TaskReward struct {
	Code     string `json:"code"`
	Quantity int    `json:"quantity"`
}

type TaskRewardData struct {
	Cooldown  Cooldown   `json:"cooldown"`
	Reward    TaskReward `json:"reward"`
	Character Character  `json:"character"`
}
