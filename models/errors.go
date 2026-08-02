package models

type ArtifactError string

func (err ArtifactError) Error() string {
	return string(err)
}

const (
	ErrBadToken       ArtifactError = "failed to parse token: token is missing or empty"
	ErrInvalidPayload ArtifactError = "request not processed: invalid payload"

	ErrMapNotFound         ArtifactError = "map not found"
	ErrItemNotFound        ArtifactError = "item not found"
	ErrCharacterNotFound   ArtifactError = "character not found"
	ErrRessourceNotFound   ArtifactError = "ressource not found on this map"
	ErrCraftNotFound       ArtifactError = "craft not found"
	ErrTaskmasterNotFound  ArtifactError = "tasks master not found on this map"
	ErrBankNotFound        ArtifactError = "bank not found"
	ErrWorkshopNotFound    ArtifactError = "workshop not found on this map"
	ErrGENotFound          ArtifactError = "grand exchange not found on this map"
	ErrAchievementNotFound ArtifactError = "achievement not found"
	ErrMonsterNotFound     ArtifactError = "monster not found"
	ErrTaskNotFound        ArtifactError = "task not found"
	ErrRewardNotFound      ArtifactError = "task reward not found"
	ErrGEOrderNotFound     ArtifactError = "order not found"

	ErrInsufficientQuantity   ArtifactError = "missing item or insufficient quantity"
	ErrTooMuchConsumables     ArtifactError = "can't equip more than 100 consumables"
	ErrItemAlreadyEquiped     ArtifactError = "item already equipped"
	ErrSlotNotEmpty           ArtifactError = "slot not empty"
	ErrLevelTooLow            ArtifactError = "level too low"
	ErrCharacterInCooldown    ArtifactError = "character in cooldown"
	ErrCharacterFullInventory ArtifactError = "character inventory is full"
	ErrActionInProgress       ArtifactError = "an action is already in progress by your character"
	ErrAlreadyAtDestination   ArtifactError = "already at destination"
	ErrInsufficientSkillLevel ArtifactError = "not skill level required"
	ErrCantBeRecycled         ArtifactError = "item can't be recycled"

	ErrCharacterAlreadyHasTask ArtifactError = "character already has a task"
	ErrCharacterHasNoTask      ArtifactError = "character has no task"
	ErrTaskNotCompleted        ArtifactError = "character has not completed the task"
	ErrTaskNotOwned            ArtifactError = "character does not have this task"
	ErrTaskAlreadyCompleted    ArtifactError = "character have already completed the task or are trying to trade too many items"

	ErrTransactionInProgress ArtifactError = "transaction already in progress with item or gold in the bank"
	ErrBankFull              ArtifactError = "bank is full"
	ErrInsufficientGold      ArtifactError = "insufficient gold"

	ErrTooManyItems          ArtifactError = "the order has not this quantity"
	ErrTransactionSelf       ArtifactError = "can't trade with yourself"
	ErrTransactionOther      ArtifactError = "transaction is already in progress on this order by a another character"
	ErrTransactionTooMany    ArtifactError = "can't create more than 100 orders"
	ErrNoStock               ArtifactError = "no stock for this item"
	ErrNoItem                ArtifactError = "no item at this price"
	ErrMissingItem           ArtifactError = "missing required item"
	ErrTransactionCharacter  ArtifactError = "transaction is already in progress by your character"
	ErrItemCannotBePurchased ArtifactError = "item cannot be purchased"
	ErrItemCannotBeSold      ArtifactError = "item cannot be sold"
	ErrNPCNotFoundOnThisMap  ArtifactError = "NPC not found on this map"
	ErrNPCNotFound           ArtifactError = "NPC not found"
)
