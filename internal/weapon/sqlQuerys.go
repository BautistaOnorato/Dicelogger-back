package weapon

var (
	QueryCreateWeapon = `INSERT INTO weapon (weapon_type, name , weight ,  price ,  category ,  reach, description , damage , versatile_damage , ammunition , damage_type, campaign_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`
    QueryGetAll = `SELECT * FROM weapon;`
	QueryGetByCampaignId = `SELECT * FROM weapon WHERE campaign_id = ? ;`
    QueryGetById = `SELECT * FROM weapon WHERE weapon_id = ? ;`
    QueryGetGeneric = `SELECT * FROM weapon WHERE campaign_id IS NULL`
    QueryUpdate = `UPDATE weapon set weapon_type = ? , name = ? , weight = ? ,  price = ? ,  category = ? ,  reach = ?, description = ? , damage = ? , versatile_damage = ? , ammunition = ? , damage_type = ? , campaign_id = ? WHERE weapon_id = ?`
    QueryDelete = `DELETE FROM weapon WHERE weapon_id = ?`
)