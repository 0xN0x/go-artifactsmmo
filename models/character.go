package models

type CharacterSkin string

const (
	CharacterSkinMen1   CharacterSkin = "men1"
	CharacterSkinMen2   CharacterSkin = "men2"
	CharacterSkinMen3   CharacterSkin = "men3"
	CharacterSkinWomen1 CharacterSkin = "women1"
	CharacterSkinWomen2 CharacterSkin = "women2"
	CharacterSkinWomen3 CharacterSkin = "women3"
)

type Character struct {
	// Main
	Name               string        `json:"name"`
	Skin               CharacterSkin `json:"skin"`
	Level              int           `json:"level"`
	Xp                 int           `json:"xp"`
	MaxXp              int           `json:"max_xp"`
	AchievementsPoints int           `json:"achievements_points"`
	Gold               int           `json:"gold"`
	Speed              int           `json:"speed"`
	Hp                 int           `json:"hp"`
	Haste              int           `json:"haste"`
	CriticalStrike     int           `json:"critical_strike"`
	Stamina            int           `json:"stamina"`

	// Position
	X int `json:"x"`
	Y int `json:"y"`

	// Mining
	MiningLevel int `json:"mining_level"`
	MiningXp    int `json:"mining_xp"`
	MiningMaxXp int `json:"mining_max_xp"`
	// Woodcutting
	WoodCuttingLevel int `json:"wood_cutting_level"`
	WoodCuttingXp    int `json:"wood_cutting_xp"`
	WoodCuttingMaxXp int `json:"wood_cutting_max_xp"`
	// Fishing
	FishingLevel int `json:"fishing_level"`
	FishingXp    int `json:"fishing_xp"`
	FishingMaxXp int `json:"fishing_max_xp"`
	// Weaponcrafting
	WeaponCraftingLevel int `json:"weaponcrafting_level"`
	WeaponCraftingXp    int `json:"weaponcrafting_xp"`
	WeaponCraftingMaxXp int `json:"weaponcrafting_max_xp"`
	// Gearcrafting
	GearCraftingLevel int `json:"gearcrafting_level"`
	GearCraftingXp    int `json:"gearcrafting_xp"`
	GearCraftingMaxXp int `json:"gearcrafting_max_xp"`
	// Jewelry
	JewelryCraftingLevel int `json:"jewelrycrafting_level"`
	JewelryCraftXp       int `json:"jewelrycrafting_xp"`
	JewelryCraftMaxXp    int `json:"jewelrycrafting_max_xp"`
	// Cooking
	CookingLevel int `json:"cooking_level"`
	CookingXp    int `json:"cooking_xp"`
	CookingMaxXp int `json:"cooking_max_xp"`

	// Attacks
	AttackFire  int `json:"attack_fire"`
	AttackEarth int `json:"attack_earth"`
	AttackWater int `json:"attack_water"`
	AttackAir   int `json:"attack_air"`

	// Damages
	DamageFire  int `json:"dmg_fire"`
	DamageEarth int `json:"dmg_earth"`
	DamageWater int `json:"dmg_water"`
	DamageAir   int `json:"dmg_air"`

	// Resistances
	ResFire  int `json:"res_fire"`
	ResEarth int `json:"res_earth"`
	ResWater int `json:"res_water"`
	ResAir   int `json:"res_air"`

	// Cooldowns
	Cooldown           int    `json:"cooldown"`
	CooldownExpiration string `json:"cooldown_expiration"`

	// Equipments Slots
	WeaponSlot              string `json:"weapon_slot"`
	ShieldSlot              string `json:"shield_slot"`
	HelmetSlot              string `json:"helmet_slot"`
	BodyArmorSlot           string `json:"body_armor_slot"`
	LegArmorSlot            string `json:"leg_armor_slot"`
	BodySlot                string `json:"body_slot"`
	Ring1Slot               string `json:"ring1_slot"`
	Ring2Slot               string `json:"ring2_slot"`
	AmuletSlot              string `json:"amulet_slot"`
	Artifact1Slot           string `json:"artifact1_slot"`
	Artifact2Slot           string `json:"artifact2_slot"`
	Artifact3Slot           string `json:"artifact3_slot"`
	Consumable1Slot         string `json:"consumable1_slot"`
	Consumable1SlotQuantity int    `json:"consumable1_slot_quantity"`
	Consumable2Slot         string `json:"consumable2_slot"`
	Consumable2SlotQuantity int    `json:"consumable2_slot_quantity"`

	// Task
	Task         string `json:"task"`
	TaskType     string `json:"task_type"`
	TaskProgress int    `json:"task_progress"`
	TaskTotal    int    `json:"task_total"`

	// Inventory
	InventoryMaxItems int              `json:"inventory_max_items"`
	Inventory         []InventorySlots `json:"inventory"`
}

type InventorySlots struct {
	Slot     int    `json:"slot"`
	Code     string `json:"code"`
	Quantity int    `json:"quantity"`
}

type Slot string

const (
	Weapon      Slot = "weapon"
	Shield      Slot = "shield"
	Helmet      Slot = "helmet"
	BodyArmor   Slot = "body_armor"
	LegArmor    Slot = "leg_armor"
	Boots       Slot = "boots"
	Ring1       Slot = "ring1"
	Ring2       Slot = "ring2"
	Amulet      Slot = "amulet"
	Artifact1   Slot = "artifact1"
	Artifact2   Slot = "artifact2"
	Artifact3   Slot = "artifact3"
	Consumable1 Slot = "consumable1"
	Consumable2 Slot = "consumable2"
)

type ItemInventory struct {
	Code     string `json:"code"`
	Slot     Slot   `json:"slot"`
	Quantity int    `json:"quantity"`
}

type RemoveItemInventory struct {
	Slot     Slot `json:"slot"`
	Quantity int  `json:"quantity"`
}
