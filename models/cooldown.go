package models

type CooldownReason string

const (
	CooldownReasonMovement     CooldownReason = "movement"
	CooldownReasonFight        CooldownReason = "fight"
	CooldownReasonCrafting     CooldownReason = "crafting"
	CooldownReasonGathering    CooldownReason = "gathering"
	CooldownReasonBuyGE        CooldownReason = "buy_ge"
	CooldownReasonSellGE       CooldownReason = "sell_ge"
	CooldownReasonDeleteItem   CooldownReason = "delete_item"
	CooldownReasonDepositBank  CooldownReason = "deposit_bank"
	CooldownReasonWithdrawBank CooldownReason = "withdraw_bank"
	CooldownReasonEquip        CooldownReason = "equip"
	CooldownReasonUnequip      CooldownReason = "unequip"
	CooldownReasonTask         CooldownReason = "task"
	CooldownReasonRecycling    CooldownReason = "recycling"
)

type Cooldown struct {
	TotalSeconds     int            `json:"total_seconds"`
	RemainingSeconds int            `json:"remaining_seconds"`
	StartedAt        string         `json:"started_at"`
	Expiration       string         `json:"expiration"`
	Reason           CooldownReason `json:"reason"`
}
