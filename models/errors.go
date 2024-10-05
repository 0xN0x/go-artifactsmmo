package models

type ArtifactError string

func (err ArtifactError) Error() string {
	return string(err)
}

const (
	ErrBadToken ArtifactError = "failed to parse token: token is missing or empty"

	ErrMapNotFound        ArtifactError = "map not found"
	ErrItemNotFound       ArtifactError = "item not found"
	ErrCharacterNotFound  ArtifactError = "character not found"
	ErrRessourceNotFound  ArtifactError = "ressource not found on this map"
	ErrCraftNotFound      ArtifactError = "craft not found"
	ErrTaskmasterNotFound ArtifactError = "tasks master not found on this map"
	ErrBankNotFound       ArtifactError = "bank not found"
	ErrWorkshopNotFound   ArtifactError = "workshop not found on this map"

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

	ErrCharacterAlreadyHasTask ArtifactError = "character already has a task"
	ErrCharacterHasNoTask      ArtifactError = "character has no task"
	ErrTaskNotCompleted        ArtifactError = "character has not completed the task"
	ErrTaskNotOwned            ArtifactError = "character does not have this task"
	ErrTaskAlreadyCompleted    ArtifactError = "character have already completed the task or are trying to trade too many items"

	ErrTransactionInProgress ArtifactError = "transaction already in progress with item or gold in the bank"
	ErrBankFull              ArtifactError = "bank is full"
	ErrInsufficientGold      ArtifactError = "insufficient gold"
)
