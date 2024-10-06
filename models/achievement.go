package models

type AchievementType string

const (
	ACombatKill  AchievementType = "combat_kill"
	ACombatDrop  AchievementType = "combat_drop"
	ACombatLevel AchievementType = "combat_level"
	AGathering   AchievementType = "gathering"
	ACrafting    AchievementType = "crafting"
	ARecycling   AchievementType = "recycling"
	ATask        AchievementType = "task"
	AOther       AchievementType = "other"
)
