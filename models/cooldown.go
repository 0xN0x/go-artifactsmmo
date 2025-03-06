package models

type CooldownReason string

const (
	CooldownReasonMovement          CooldownReason = "movement"
	CooldownReasonFight             CooldownReason = "fight"
	CooldownReasonCrafting          CooldownReason = "crafting"
	CooldownReasonGathering         CooldownReason = "gathering"
	CooldownReasonBuyGE             CooldownReason = "buy_ge"
	CooldownReasonSellGE            CooldownReason = "sell_ge"
	CooldownReasonBuyNPC            CooldownReason = "buy_npc"
	CooldownReasonSellNPC           CooldownReason = "sell_npc"
	CooldownReasonCancelGE          CooldownReason = "cancel_ge"
	CooldownReasonDeleteItem        CooldownReason = "delete_item"
	CooldownReasonDeposit           CooldownReason = "deposit"
	CooldownReasonWithdraw          CooldownReason = "withdraw"
	CooldownReasonDepositGold       CooldownReason = "deposit_gold"
	CooldownReasonWithdrawGold      CooldownReason = "withdraw_gold"
	CooldownReasonEquip             CooldownReason = "equip"
	CooldownReasonUnequip           CooldownReason = "unequip"
	CooldownReasonTask              CooldownReason = "task"
	CooldownReasonChristmasExchange CooldownReason = "christmas_exchange"
	CooldownReasonRecycling         CooldownReason = "recycling"
	CooldownReasonRest              CooldownReason = "rest"
	CooldownReasonUse               CooldownReason = "use"
	CooldownReasonBuyBankExpansion  CooldownReason = "buy_bank_expansion"
)

type Cooldown struct {
	TotalSeconds     int            `json:"total_seconds"`
	RemainingSeconds int            `json:"remaining_seconds"`
	StartedAt        string         `json:"started_at"`
	Expiration       string         `json:"expiration"`
	Reason           CooldownReason `json:"reason"`
}
