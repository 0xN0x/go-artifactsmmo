package models

type ArtifactError string

func (err ArtifactError) Error() string {
	return string(err)
}

const (
	ErrBadToken ArtifactError = "failed to parse token: token is missing or empty"

	ErrMapNotFound       ArtifactError = "map not found"
	ErrItemNotFound      ArtifactError = "item not found"
	ErrCharacterNotFound ArtifactError = "character not found"
	ErrRessourceNotFound ArtifactError = "ressource not found on this map"

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
)