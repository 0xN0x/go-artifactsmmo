package models

type CooldownReason string

const (
	CooldownReasonMovement         CooldownReason = "movement"
	CooldownReasonFight            CooldownReason = "fight"
	CooldownReasonMultiFight       CooldownReason = "multi_fight"
	CooldownReasonCrafting         CooldownReason = "crafting"
	CooldownReasonGathering        CooldownReason = "gathering"
	CooldownReasonBuyGE            CooldownReason = "buy_ge"
	CooldownReasonSellGE           CooldownReason = "sell_ge"
	CooldownReasonCreateGE         CooldownReason = "create_buy_order_ge"
	CooldownReasonFillBuyGE        CooldownReason = "fill_buy_order_ge"
	CooldownReasonBuyNPC           CooldownReason = "buy_npc"
	CooldownReasonSellNPC          CooldownReason = "sell_npc"
	CooldownReasonCancelGE         CooldownReason = "cancel_ge"
	CooldownReasonDeleteItem       CooldownReason = "delete_item"
	CooldownReasonDepositItem      CooldownReason = "deposit_item"
	CooldownReasonWithdrawItem     CooldownReason = "withdraw_item"
	CooldownReasonDepositGold      CooldownReason = "deposit_gold"
	CooldownReasonWithdrawGold     CooldownReason = "withdraw_gold"
	CooldownReasonEquip            CooldownReason = "equip"
	CooldownReasonUnequip          CooldownReason = "unequip"
	CooldownReasonTask             CooldownReason = "task"
	CooldownReasonRecycling        CooldownReason = "recycling"
	CooldownReasonRest             CooldownReason = "rest"
	CooldownReasonUse              CooldownReason = "use"
	CooldownReasonBuyBankExpansion CooldownReason = "buy_bank_expansion"
	CooldownReasonGiveItem         CooldownReason = "give_item"
	CooldownReasonGiveGold         CooldownReason = "give_gold"
	CooldownReasonChangeSkin       CooldownReason = "change_skin"
	CooldownReasonRename           CooldownReason = "rename"
	CooldownReasonTransition       CooldownReason = "transition"
	CooldownReasonClaimItem        CooldownReason = "claim_item"
	CooldownReasonSandboxGiveGold  CooldownReason = "sandbox_give_gold"
	CooldownReasonSandboxGiveItem  CooldownReason = "sandbox_give_item"
	CooldownReasonSandboxGiveXp    CooldownReason = "sandbox_give_xp"
)

type Cooldown struct {
	TotalSeconds     int            `json:"total_seconds"`
	RemainingSeconds int            `json:"remaining_seconds"`
	StartedAt        string         `json:"started_at"`
	Expiration       string         `json:"expiration"`
	Reason           CooldownReason `json:"reason"`
}
