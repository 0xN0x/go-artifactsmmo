package models

type CooldownReason string

const (
	CooldownReasonMovement             CooldownReason = "movement"
	CooldownReasonFight                CooldownReason = "fight"
	CooldownReasonRaidFight            CooldownReason = "raid_fight"
	CooldownReasonMultiFight           CooldownReason = "multi_fight"
	CooldownReasonCrafting             CooldownReason = "crafting"
	CooldownReasonGathering            CooldownReason = "gathering"
	CooldownReasonCreateBuyOrderGE     CooldownReason = "create_buy_order_ge"
	CooldownReasonFillBuyOrderGE       CooldownReason = "fill_buy_order_ge"
	CooldownReasonBuyNPC               CooldownReason = "buy_npc"
	CooldownReasonSellNPC              CooldownReason = "sell_npc"
	CooldownReasonCancelGE             CooldownReason = "cancel_ge"
	CooldownReasonDeleteItem           CooldownReason = "delete_item"
	CooldownReasonDepositItem          CooldownReason = "deposit_item"
	CooldownReasonWithdrawItem         CooldownReason = "withdraw_item"
	CooldownReasonDepositGold          CooldownReason = "deposit_gold"
	CooldownReasonWithdrawGold         CooldownReason = "withdraw_gold"
	CooldownReasonEquip                CooldownReason = "equip"
	CooldownReasonUnequip              CooldownReason = "unequip"
	CooldownReasonTask                 CooldownReason = "task"
	CooldownReasonRecycling            CooldownReason = "recycling"
	CooldownReasonRest                 CooldownReason = "rest"
	CooldownReasonUse                  CooldownReason = "use"
	CooldownReasonBuyBankExpansion     CooldownReason = "buy_bank_expansion"
	CooldownReasonGiveItem             CooldownReason = "give_item"
	CooldownReasonGiveGold             CooldownReason = "give_gold"
	CooldownReasonRaidDeposit          CooldownReason = "raid_deposit"
	CooldownReasonChangeSkin           CooldownReason = "change_skin"
	CooldownReasonRename               CooldownReason = "rename"
	CooldownReasonTransition           CooldownReason = "transition"
	CooldownReasonClaimItem            CooldownReason = "claim_item"
	CooldownReasonSandboxGiveGold      CooldownReason = "sandbox_give_gold"
	CooldownReasonSandboxGiveItem      CooldownReason = "sandbox_give_item"
	CooldownReasonSandboxGiveXp        CooldownReason = "sandbox_give_xp"
	CooldownReasonSandboxClearCooldown CooldownReason = "sandbox_clear_cooldown"
	CooldownReasonSandboxTeleport      CooldownReason = "sandbox_teleport"
	//CooldownReasonChristmasExchange CooldownReason = "christmas_exchange"
)

type Cooldown struct {
	TotalSeconds     int            `json:"total_seconds"`
	RemainingSeconds int            `json:"remaining_seconds"`
	StartedAt        string         `json:"started_at"`
	Expiration       string         `json:"expiration"`
	Reason           CooldownReason `json:"reason"`
}
