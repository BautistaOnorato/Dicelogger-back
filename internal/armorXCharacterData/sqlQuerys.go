package armorXCharacterData

var (
	QueryCreateCharacterArmor  = `INSERT INTO character_armor (character_id, armor_id, equipped) VALUES (?, ?, ?)`
	QueryGetAllCharacterArmor  = `SELECT character_armor.character_armor_id, character_id, armor.armor_id, armor.material, armor.name, armor.weight, armor.price, armor.category, armor.protection_type, armor.description, armor.penalty, armor.strength, armor.armor_class, armor.dex_bonus, armor.campaign_id, equipped FROM character_armor left join armor on character_armor.armor_id = armor.armor_id`
	QueryGetByIdCharacterArmor = `SELECT character_armor.character_armor_id, character_id, armor.armor_id, armor.material, armor.name, armor.weight, armor.price, armor.category, armor.protection_type, armor.description, armor.penalty, armor.strength, armor.armor_class, armor.dex_bonus, armor.campaign_id, equipped FROM character_armor left join armor on character_armor.armor_id = armor.armor_id WHERE character_armor_id = ?`
	QueryGetByCharacterId      = `SELECT character_armor.character_armor_id, character_id, armor.armor_id, armor.material, armor.name, armor.weight, armor.price, armor.category, armor.protection_type, armor.description, armor.penalty, armor.strength, armor.armor_class, armor.dex_bonus, armor.campaign_id, equipped FROM character_armor left join armor on character_armor.armor_id = armor.armor_id WHERE character_id = ?`
	QueryUpdateCharacterArmor  = `UPDATE character_armor SET character_id = ?, armor_id = ?, equipped = ? WHERE character_armor_id = ?`
	QueryUpdateOwnership  = `UPDATE character_armor SET character_id = ?, equipped = ? WHERE character_armor_id = ?`
	QueryDeleteCharacterArmor  = `DELETE FROM character_armor WHERE character_armor_id = ?`
	QueryDeleteByCharacterId   = `DELETE FROM character_armor WHERE character_id = ?`
)
