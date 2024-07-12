package characterxattackevent

var (
	QueryGetAll = `SELECT * FROM character_attack_event;`
	QueryGetById = `SELECT * FROM character_attack_event WHERE character_event=?;`
	QueryGetByCharacterId = `SELECT * FROM character_attack_event WHERE character_id=?;`
	QueryGetByEventId = `SELECT * FROM character_attack_event WHERE event_id=?;`
	QueryInsert = `INSERT INTO character_attack_event (event_id, character_id, dmg, dmg_roll, attack_result, attack_roll, armor_class) VALUES (?, ?, ?, ?, ?, ?, ?);`
	QueryDelete = `DELETE FROM character_attack_event WHERE character_event=?;`
)