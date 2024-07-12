package weaponxcharacterdata

var (
	QueryCreateWeaponXCharacterData = `INSERT INTO character_weapon (character_id, weapon_id, equipped) VALUES (?, ?, ?);`
    QueryGetAll = `SELECT character_weapon_id, character_id, weapon.weapon_id, weapon_type, name, weight, price, category, reach, description, damage, versatile_damage, ammunition, damage_type, campaign_id, equipped FROM character_weapon LEFT JOIN weapon ON character_weapon.weapon_id = weapon.weapon_id;`
    QueryGetById = `SELECT character_weapon_id, character_id, weapon.weapon_id, weapon_type, name, weight, price, category, reach, description, damage, versatile_damage, ammunition, damage_type, campaign_id, equipped FROM character_weapon LEFT JOIN weapon ON character_weapon.weapon_id = weapon.weapon_id WHERE character_weapon_id = ? ;`
    QueryGetByCharacterDataId = `SELECT character_weapon_id, character_id, weapon.weapon_id, weapon_type, name, weight, price, category, reach, description, damage, versatile_damage, ammunition, damage_type, campaign_id, equipped FROM character_weapon LEFT JOIN weapon ON character_weapon.weapon_id = weapon.weapon_id WHERE character_id = ?;`
    QueryUpdate= `UPDATE character_weapon SET character_id = ?, weapon_id = ?, equipped = ? where character_weapon_id = ?`
    QueryUpdateOwnership= `UPDATE character_weapon SET character_id = ?, equipped = ? where character_weapon_id = ?`
    QueryDelete= `DELETE FROM character_weapon WHERE character_weapon_id = ?;`
    QueryDeleteByCharacterDataId = `DELETE FROM character_weapon WHERE character_id = ?;`
)